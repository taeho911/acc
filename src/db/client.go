package db

import (
	"os"
	"sync"
	"time"
	"context"
	"strings"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

const usernameKey string = "MONGO_USERNAME"
const passwordKey string = "MONGO_PASSWORD"
const host string = "localhost"
const port string = "27017"
const dbserver string = "mongodb://" + host + ":" + port
const DATABASE string = "taeho"
const COLLECTION string = "acc"

var once sync.Once
var client *mongo.Client
var ctx context.Context
var cancel context.CancelFunc

type Acc struct {
	Index int
	Title string
	Url string
	Uid string
	Pwd string
	Email string
	Alias []string
	Memo string
}

func NewClient() {
	once.Do(func() {
		username := os.Getenv(usernameKey)
		password := os.Getenv(passwordKey)

		var sb strings.Builder
		sb.WriteString("mongodb://")
		sb.WriteString(username)
		sb.WriteString(":")
		sb.WriteString(password)
		sb.WriteString("@")
		sb.WriteString(host)
		sb.WriteString(":")
		sb.WriteString(port)
		sb.WriteString("/?authSource=admin")

		connectionURL := sb.String()

		var err error
		client, err = mongo.NewClient(options.Client().ApplyURI(connectionURL))
		if err != nil {
			log.Panicln("Error: Failed to create database client")
		}
		ctx, cancel = context.WithTimeout(context.Background(), 5 * time.Second)
		// defer cancel()

		err = client.Connect(ctx)
		if err != nil {
			log.Panicln("Error: Failed to connect to database")
		}
		// defer client.Disconnect(ctx)
	})
}

func DelClient() {
	if client != nil && cancel != nil {
		client.Disconnect(ctx)
		cancel()
	}
}

func SetUsername(username string) {
	err := os.Setenv(usernameKey, username)
	if err != nil {
		log.Panicln("Error: Failed to set username")
	}
}

func SetPassword(password string) {
	err := os.Setenv(passwordKey, password)
	if err != nil {
		log.Panicln("Error: Failed to set password")
	}
}

func UnsetUsername() {
	err := os.Unsetenv(usernameKey)
	if err != nil {
		log.Panicln("Error: Failed to unset username")
	}
}

func UnsetPassword() {
	err := os.Unsetenv(passwordKey)
	if err != nil {
		log.Panicln("Error: Failed to unset password")
	}
}

func PingConnection() {
	if client == nil {
		log.Println("No client for database")
	} else {
		err := client.Ping(ctx, readpref.Primary())
		if err != nil {
			log.Println("Failed to authenticate")
			log.Println(err)
		} else {
			log.Println("Succeed to connect to", dbserver)
		}
	}
}

func setAndGetNextSeq(coll *mongo.Collection) interface{} {
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

func Test() {
	NewClient()
	coll := client.Database(DATABASE).Collection(COLLECTION)
	seqNum := setAndGetNextSeq(coll)
	log.Println("updatedDoc[\"seq_num\"]=", seqNum)
}