package db

import (
	"os"
	"time"
	"context"
	"strings"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

const usernameKey string = "MONGO_USERNAME"
const passwordKey string = "MONGO_PASSWORD"
const database string = "taeho"
const collection string = "acc"

type Account struct {
	Id primitive.ObjectID `bson:"_id",omitempty`
	Title string `bson:"title",omitempty`
	Url string `bson:"url",omitempty`
	Uid string `bson:"uid",omitempty`
	Pwd string `bson:"pwd",omitempty`
	Email string `bson:"email",omitempty`
	Alias []string `bson:"alias",omitempty`
	Memo string `bson:"memo",omitempty`
}

// INTERNAL FUNCTIONS ============================================================

func getClient() (*mongo.Client, context.Context, context.CancelFunc) {
	username := os.Getenv(usernameKey)
	password := os.Getenv(passwordKey)

	var sb strings.Builder
	sb.WriteString("mongodb://")
	sb.WriteString(username)
	sb.WriteString(":")
	sb.WriteString(password)
	sb.WriteString("@localhost:27017/?authSource=admin")
	connectionURL := sb.String()

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURL))
	if err != nil {
		log.Panicln("Error: Failed to create database client")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Panicln("Error: Failed to connect to database")
	}
	return client, ctx, cancel
}

func delClient(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	client.Disconnect(ctx)
	cancel()
}

func issueNextSeq(coll *mongo.Collection, ctx context.Context) interface{} {
	var updatedDoc bson.M
	filter := bson.D{{"_id", "seq"}}
	update := bson.D{{"$inc", bson.D{{"seq_num", 1}}}}
	option := options.FindOneAndUpdate().SetUpsert(true)
	err := coll.FindOneAndUpdate(ctx, filter, update, option).Decode(&updatedDoc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("There was no document for _id=seq")
			return nil
		}
		log.Panicln(err)
	}
	log.Println("updated document", updatedDoc)
	return updatedDoc["seq_num"]
}

// EXPORT FUNCTIONS ============================================================

func Ping() bool {
	client, ctx, cancel := getClient()
	defer delClient(client, ctx, cancel)

	if client == nil {
		log.Println("No client for database")
		return false
	} else {
		err := client.Ping(ctx, readpref.Primary())
		if err != nil {
			log.Println("Failed to authenticate")
			log.Println(err)
			return false
		} else {
			log.Println("Succeed to connect")
			return true
		}
	}
}

func InsertOne(title, url, uid, pwd, email, memo string, aliasArr []string) *mongo.InsertOneResult {
	client, ctx, cancel := getClient()
	defer delClient(client, ctx, cancel)
	coll := client.Database(database).Collection(collection)

	document := bson.M{
		"_id": issueNextSeq(coll, ctx),
		"title": title,
		"url": url,
		"uid": uid,
		"pwd": pwd,
		"email": email,
		"alias": aliasArr,
		"memo": memo,
	}

	result, err := coll.InsertOne(ctx, document)
	if err != nil {
		log.Println("Failed to insert")
		return nil
	}

	return result
}

func DeleteMany(idArr []int) *mongo.DeleteResult {
	client, ctx, cancel := getClient()
	defer delClient(client, ctx, cancel)
	coll := client.Database(database).Collection(collection)

	document := bson.M{
		"_id": bson.D{{"$in", idArr}},
	}

	result, err := coll.DeleteMany(ctx, document)
	if err != nil {
		log.Println("Failed to delete")
		return nil
	}

	return result
}

func FindAll() []bson.M {
	client, ctx, cancel := getClient()
	defer delClient(client, ctx, cancel)
	coll := client.Database(database).Collection(collection)

	document := bson.M{
		"_id": bson.D{{"$not", bson.D{{"$eq", "seq"}}}},
	}
	opts := options.Find().SetSort(bson.D{{"_id", 1}})

	cursor, err := coll.Find(ctx, document, opts)
	if err != nil {
		log.Println(err)
		log.Panic("Failed to find all - 1")
		return nil
	}
	var result []bson.M
	if err = cursor.All(ctx, &result); err != nil {
		log.Println(err)
		log.Panic("Failed to find all - 2")
		return nil
	}

	return result
}

func Find(id int, title, alias, uid string) []Account {
	client, ctx, cancel := getClient()
	defer delClient(client, ctx, cancel)
	coll := client.Database(database).Collection(collection)

	document := bson.M{
		"_id": id,
		"title": bson.D{{"$regex", title}},
		"alias": bson.D{{"$all", alias}},
		"uid": bson.D{{"$regex", uid}},
	}
	opts := options.Find().SetSort(bson.D{{"_id", 1}})
	
	cursor, err := coll.Find(ctx, document, opts)
	if err != nil {
		log.Println(err)
		log.Panic("Failed to find - 1")
		return nil
	}
	var result []Account
	if err = cursor.All(ctx, &result); err != nil {
		log.Println(err)
		log.Panic("Failed to find - 2")
		return nil
	}

	return result
}

// TEST FUNCTIONS ============================================================