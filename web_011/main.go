package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
	"time"
)

type Member struct {
	ID 		primitive.ObjectID 	`bson:"_id,omitempty"`
	Name 	string 				`bson:"name,omitempty"`
	Author 	string 				`bson:"author,omitempty"`
	tags	[]string			`bson:"tags,omitempty"`
}

type Information struct {
	ID				primitive.ObjectID 	`bson:"_id,omitempty"`
	Member 			primitive.ObjectID 	`bson:"member,omitempty"`
	Title			string				`bson:"title,omitempty"`
	Description		string				`bson:"description,omitempty"`
	Duration		string				`bson:"duration,omitempty"`
}

func main() {

	//db 연결
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://gothic:1234@cluster0-shard-00-00.5o8nt.mongodb.net:27017,cluster0-shard-00-01.5o8nt.mongodb.net:27017,cluster0-shard-00-02.5o8nt.mongodb.net:27017/tritech?ssl=true&replicaSet=atlas-8qx0b5-shard-0&authSource=admin&retryWrites=true&w=majority"))
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
	//tritechstartdatabase := client.Database("tritech")
	//membersCollection := tritechstartdatabase.Collection("members")
	//informationsCollection := tritechstartdatabase.Collection("informations")

	//https://www.youtube.com/watch?v=uvSzIZlh8O8 9:15

	//mongoMember := Member{
	//	Name: "The MongoDB Member",
	//	Author: "",
	//}

	//기존 특정 테이블 할당 1
	//cursor, err := memberCollection.Find(ctx, bson.M{})
	//if err != nil {
	//	log.Fatal(err)
	//}

	//기존 특정 테이블 데이터 대입
	//var member []bson.M
	//if err = cursor.All(ctx, &member); err != nil {
	//	log.Fatal(err)
	//}

	//for _, members := range member {
	//	fmt.Println(members["_id"])
	//}

	//데이터 입력 방식 선언 및 입력 - 하나의 데이터만 추가
	//memberResult, err := membersCollection.InsertOne(ctx, bson.D{
	//	{"title", "The Prolyglot Developer Podcast"},
	//	{"author", "Nic Raboy"},
	//	{"tags", bson.A{"development", "programming", "coding"}},
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(memberResult.InsertedID)

	//데이터 입력 방식 선언 및 입력 - 여러 데이터 추가
	//informationResult, err := informationsCollection.InsertMany(ctx, []interface{}{
	//	bson.D{
	//		{"member", memberResult.InsertedID},
	//		{"title", "Episode #1"},
	//		{"description", "this is the first episode"},
	//		{"duration", "25"},
	//	},
	//	bson.D{
	//		{"member", memberResult.InsertedID},
	//		{"title", "Episode #2"},
	//		{"description", "this is the second episode"},
	//		{"duration", "32"},
	//	},
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(informationResult.InsertedIDs)

	//데이터 검색용
	//var members []bson.M
	//memberCursor, err := membersCollection.Find(ctx, bson.M{})
	//if err != nil {
	//	panic(err)
	//}
	//if err = memberCursor.All(ctx, &members); err != nil {
	//	panic(err)
	//}
	//fmt.Println(members)
	//
	//var informations []bson.M
	//informationCursor, err := informationsCollection.Find(ctx, bson.M{})
	//if err != nil {
	//	panic(err)
	//}
	//if err = informationCursor.All(ctx, &informations); err != nil {
	//	panic(err)
	//}
	//fmt.Println(informations)

	//데이터 형 지정 후 데이터 검색
	//var members []Member
	//memberCursor, err := membersCollection.Find(ctx, bson.M{})
	//if err != nil {
	//	panic(err)
	//}
	//if err = memberCursor.All(ctx, &members); err != nil {
	//	panic(err)
	//}
	//fmt.Println(members)
	//
	//var informations []Information
	//informationCursor, err := informationsCollection.Find(ctx, bson.M{})
	//if err != nil {
	//	panic(err)
	//}
	//if err = informationCursor.All(ctx, &informations); err != nil {
	//	panic(err)
	//}
	//fmt.Println(informations)

	//DB 및 테이블 정의
	//database := client.Database("tritech")
	//membersCollection := database.Collection("members")
	//informationsCollection := database.Collection("informations")

	//단일 인스턴스 제거
	//result, err := informationsCollection.DeleteOne(ctx, bson.M{"duration": "25"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)

	//다중 인스턴스 제거
	//result, err := informationsCollection.DeleteMany(ctx, bson.M{"duration": "32"})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("DeleteMany removed %v document(s)", result.DeletedCount)

	//테이블 데이터 제거
	//if err = membersCollection.Drop(ctx); err != nil {
	//	log.Fatal(err)
	//}

	//DB & Table Definition
	//database := client.Database("tritech")
	//informationsCollection := database.Collection("informations")
	//
	//id 정의
	//id, _ := primitive.ObjectIDFromHex("5f54be11681ae895406e866f")

	//파라미터를 통한 property update
	//result, err := informationsCollection.UpdateOne(
	//	ctx,
	//	bson.M{"_id": id},
	//	bson.D{
	//		{"$set", bson.D{{"author", "Nicolas Raboy"}}},
	//	},
	//)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("UpdateOne %v document(s)\n", result.ModifiedCount)

	//DB에 해당되는 스트링값을 통한 속성값 수정
	//result, err = informationsCollection.UpdateMany(
	//	ctx,
	//	bson.M{"title": "Sin and Punishment"},
	//	bson.D{
	//		{"$set", bson.D{{"author", "Gothic"}}},
	//	},
	//)
	//fmt.Printf("UpdateOne %v document(s)\n", result.ModifiedCount)

	//DB에 해당되는 값 대체
	//result, err = informationsCollection.ReplaceOne(
	//	ctx,
	//	bson.M{"author": "Nicolas Raboy"},
	//	bson.M{
	//		"title": "The Nic Raboy Story",
	//		"author": nil,
	//	},
	//)
	//fmt.Printf("UpdateOne %v document(s)\n", result.ModifiedCount)
}
