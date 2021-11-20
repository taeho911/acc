package api

import (
	"context"
	"errors"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB client를 생성, 접속하는 함수
func getClient() (*mongo.Client, context.Context, context.CancelFunc, error) {
	username := os.Getenv(usernameKey)
	password := os.Getenv(passwordKey)
	protocol := os.Getenv(dbProtocolKey)
	host := os.Getenv(dbHostKey)
	port := os.Getenv(dbPortKey)

	if protocol == "" || host == "" || port == "" {
		return nil, nil, nil, errors.New("lack of env")
	}

	// URL pattern when there is auth
	// mongodb://username:password@host:port/?authSource=admin
	var sb strings.Builder
	sb.WriteString(protocol)
	sb.WriteString("://")
	if username != "" && password != "" {
		sb.WriteString(username)
		sb.WriteString(":")
		sb.WriteString(password)
		sb.WriteString("@")
	}
	sb.WriteString(host)
	sb.WriteString(":")
	sb.WriteString(port)
	if username != "" && password != "" {
		sb.WriteString("/?authSource=admin")
	}
	connectionURL := sb.String()

	client, err1 := mongo.NewClient(options.Client().ApplyURI(connectionURL))
	if err1 != nil {
		return nil, nil, nil, err1
	}
	// 빈 context에 5초의 타임아웃을 추가해 DB 작업용 context를 생성
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	err2 := client.Connect(ctx)
	if err2 != nil {
		return nil, nil, nil, err2
	}

	return client, ctx, cancel, nil
}

func delClient(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	client.Disconnect(ctx)
	cancel()
}

// MongoDB는 RDBMS와는 다르게 시퀜스 객체가 없으므로, 각 collection 이름별로 시퀜스를 생성시,
// 필요에 따라 시퀘스는 인크리멘트 시키고 시퀜스 번호를 반환하는 함수
// 기본적으로 INSERT시에 사용된다.
func getNextSeq(client *mongo.Client, ctx context.Context, dbName string, collName string) (int, error) {
	coll := client.Database(dbName).Collection("seq")
	filter := bson.D{{"_id", collName}}
	update := bson.D{{"$inc", bson.D{{"seq_num", 1}}}}
	opts := options.FindOneAndUpdate().SetUpsert(true)

	var updatedDoc bson.M
	err1 := coll.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedDoc)
	if err1 != nil {
		return -1, err1
	}

	return int(updatedDoc["seq_num"].(float64)), nil
}

// INSERT MANY를 할 경우 매번 시퀜스를 인크리멘트하는 것은 성능적으로 좋지 않으므로 한꺼번에 인크리멘트하는 함수
func getNextSeqMany(client *mongo.Client, ctx context.Context, dbName string, collName string, inc int) ([]int, error) {
	coll := client.Database(dbName).Collection("seq")
	filter := bson.D{{"_id", collName}}
	update := bson.D{{"$inc", bson.D{{"seq_num", inc}}}}
	opts := options.FindOneAndUpdate().SetUpsert(true)

	var updatedDoc bson.M
	err1 := coll.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedDoc)
	if err1 != nil {
		return nil, err1
	}

	var result []int
	startNum := int(updatedDoc["seq_num"].(float64))
	for i := 0; i < inc; i++ {
		result = append(result, startNum+i)
	}

	return result, nil
}
