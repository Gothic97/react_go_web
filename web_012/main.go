package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/net/context"
	"log"
	"time"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://gothic:1234@cluster0-shard-00-00.5o8nt.mongodb.net:27017,cluster0-shard-00-01.5o8nt.mongodb.net:27017,cluster0-shard-00-02.5o8nt.mongodb.net:27017/tritech?ssl=true&replicaSet=atlas-8qx0b5-shard-0&authSource=admin&retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
}
