package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
	"time"
)

func main() {

	//db 연결
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://gothic:1234@cluster0.5o8nt.mongodb.net/tritech?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	//context 작성
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//db 연결 끊기
	defer client.Disconnect(ctx)

	//기존 특정 db 및 해당 테이블 선언
	tritechstartdatabase := client.Database("tritech")
	memberCollection := tritechstartdatabase.Collection("member")

	//기존 특정 테이블 할당 1
	cursor, err := memberCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	//기존 특정 테이블 데이터 대입
	var member []bson.M
	if err = cursor.All(ctx, &member); err != nil {
		log.Fatal(err)
	}

	for _, members := range member {
		fmt.Println(members["_id"])
	}

}
