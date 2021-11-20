package api

import (
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var dbName string = os.Getenv(dbNameKey)
var collName string = os.Getenv(collNameKey)

// DB에 접근 가능한지 판별하는 함수
func Ping() (bool, error) {
	client, ctx, cancel, err1 := getClient()
	defer delClient(client, ctx, cancel)
	if err1 != nil {
		return false, err1
	}
	err2 := client.Ping(ctx, readpref.Primary())
	if err2 != nil {
		return false, err2
	}
	return true, nil
}

func InsertOne(acc Acc) (interface{}, error) {
	client, ctx, cancel, err1 := getClient()
	defer delClient(client, ctx, cancel)
	if err1 != nil {
		return nil, err1
	}
	seqNum, err2 := getNextSeq(client, ctx, dbName, collName)
	if err2 != nil {
		return nil, err2
	}
	coll := client.Database(dbName).Collection(collName)
	acc.Index = seqNum
	acc.Deleted = false
	result, err3 := coll.InsertOne(ctx, acc)
	if err3 != nil {
		return nil, err3
	}
	return result.InsertedID, nil
}

func InsertMany(accs []Acc) ([]interface{}, error) {
	client, ctx, cancel, err1 := getClient()
	defer delClient(client, ctx, cancel)
	if err1 != nil {
		return nil, err1
	}
	seqNums, err2 := getNextSeqMany(client, ctx, dbName, collName, len(accs))
	if err2 != nil {
		return nil, err2
	}
	var docs bson.A
	for i := 0; i < len(seqNums); i++ {
		accs[i].Index = seqNums[i]
		accs[i].Deleted = false
		docs = append(docs, accs[i])
	}
	coll := client.Database(dbName).Collection(collName)
	result, err3 := coll.InsertMany(ctx, docs)
	if err3 != nil {
		return nil, err3
	}
	return result.InsertedIDs, nil
}

func DeleteMany(indexSlice []int) (int, error) {
	client, ctx, cancel, err1 := getClient()
	defer delClient(client, ctx, cancel)
	if err1 != nil {
		return -1, err1
	}
	coll := client.Database(dbName).Collection(collName)
	filter := bson.M{"index": bson.M{"$in": indexSlice}}
	update := bson.M{"$set": bson.M{"deleted": true}}
	result, err2 := coll.UpdateMany(ctx, filter, update)
	if err2 != nil {
		return -1, err2
	}
	return int(result.ModifiedCount), nil
}

// 기존의 데이터에서 특정 필드(컬럼)만을 비우고 싶을때 사용된다.
// del 서브커맨드에서 사용되는 함수
func EmptyFields(indexSlice []int, title, username, password, location, email, memo, alias bool) (int, error) {
	client, ctx, cancel, err1 := getClient()
	defer delClient(client, ctx, cancel)
	if err1 != nil {
		return -1, err1
	}
	coll := client.Database(dbName).Collection(collName)
	filter := bson.M{"index": bson.D{{"$in", indexSlice}}}
	setMap := bson.M{}
	if title {
		setMap["title"] = ""
	}
	if username {
		setMap["username"] = ""
	}
	if password {
		setMap["password"] = ""
	}
	if location {
		setMap["location"] = ""
	}
	if email {
		setMap["email"] = ""
	}
	if memo {
		setMap["memo"] = ""
	}
	if alias {
		setMap["alias"] = []string{}
	}
	update := bson.M{"$set": setMap}
	result, err2 := coll.UpdateMany(ctx, filter, update)
	if err2 != nil {
		return -1, err2
	}
	return int(result.ModifiedCount), nil
}

func FindAll() ([]Acc, error) {
	client, ctx, cancel, err1 := getClient()
	defer delClient(client, ctx, cancel)
	if err1 != nil {
		return nil, err1
	}
	coll := client.Database(dbName).Collection(collName)
	filter := bson.M{"deleted": false}
	opts := options.Find().SetSort(bson.D{{"index", 1}})
	cursor, err2 := coll.Find(ctx, filter, opts)
	if err2 != nil {
		return nil, err2
	}
	defer cursor.Close(ctx)
	var result []Acc
	if err3 := cursor.All(ctx, &result); err3 != nil {
		return nil, err3
	}
	return result, nil
}

func Find(index int, title, username string, aliasSlice []string) ([]Acc, error) {
	client, ctx, cancel, err1 := getClient()
	defer delClient(client, ctx, cancel)
	if err1 != nil {
		return nil, err1
	}
	coll := client.Database(dbName).Collection(collName)
	var conditions []bson.M
	conditions = append(conditions, bson.M{"deleted": false})
	if index != 0 {
		conditions = append(conditions, bson.M{"index": index})
	}
	if title != "" {
		conditions = append(conditions, bson.M{"title": bson.D{{"$regex", title}}})
	}
	if username != "" {
		conditions = append(conditions, bson.M{"username": bson.D{{"$regex", username}}})
	}
	if len(aliasSlice) > 0 {
		conditions = append(conditions, bson.M{"alias": bson.D{{"$in", aliasSlice}}})
	}
	filter := bson.M{"$and": conditions}
	opts := options.Find().SetSort(bson.D{{"index", 1}})
	cursor, err2 := coll.Find(ctx, filter, opts)
	if err2 != nil {
		return nil, err2
	}
	defer cursor.Close(ctx)
	var result []Acc
	if err3 := cursor.All(ctx, &result); err3 != nil {
		return nil, err3
	}
	return result, nil
}

func UpdateMany(indexSlice []int, acc Acc, aliasPull, aliasPush bool) (int, error) {
	client, ctx, cancel, err1 := getClient()
	if err1 != nil {
		return -1, err1
	}
	defer delClient(client, ctx, cancel)
	coll := client.Database(dbName).Collection(collName)
	filter := bson.M{"index": bson.D{{"$in", indexSlice}}}
	setMap := bson.M{}
	if acc.Title != "" {
		setMap["title"] = acc.Title
	}
	if acc.Username != "" {
		setMap["username"] = acc.Username
	}
	if acc.Password != "" {
		setMap["password"] = acc.Password
	}
	if acc.Location != "" {
		setMap["location"] = acc.Location
	}
	if acc.Email != "" {
		setMap["email"] = acc.Email
	}
	if len(acc.Alias) > 0 {
		if aliasPush {
			setMap["$push"] = bson.D{{"alias", bson.D{{"$each", acc.Alias}}}}
		} else if aliasPull {
			setMap["$pull"] = bson.D{{"alias", bson.D{{"$each", acc.Alias}}}}
		} else {
			setMap["alias"] = acc.Alias
		}
	}
	if acc.Memo != "" {
		setMap["memo"] = acc.Memo
	}
	update := bson.M{"$set": setMap}
	result, err2 := coll.UpdateMany(ctx, filter, update)
	if err2 != nil {
		return -1, err2
	}
	return int(result.ModifiedCount), nil
}
