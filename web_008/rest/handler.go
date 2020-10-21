package rest

import (
	_ "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
	"net/http"
	_ "sort"
	"time"
)

type Handler struct {
	test int64
}

func MakeHandler() (*Handler, error) {
	return new(Handler), nil
}

func (h *Handler) GetMainPage(c *gin.Context) {
	log.Println("Main page....")
	c.String(http.StatusOK, "Main page for secure API!!")
	//fmt.Fprintf(c.Writer, "Main page for secure API!!")
}

func (h *Handler) GetTest(c *gin.Context) {
	log.Println("Test")
	c.String(http.StatusOK, "Test page for secure API!!")
	fmt.Fprintf(c.Writer, "TEST")
}

func (h *Handler) GetProducts(c *gin.Context) {
	fmt.Fprintf(c.Writer, "Hi I'm GET")
}

//type CRUD struct {
//	dbname 		string				`json:  "dbname"`
//	tbname		string				`json:  "tbname"`
//	create		interface{}			`json:	"create"`
//	read		interface{}			`json:	"read"`
//	update		interface{}			`json:	"update"`
//	delete		interface{}			`json:  "delete"`
//	created_at	time.Time			`json:  "created_at"`
//}

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

//type Post struct {
//
//	Title string `json:”title,omitempty”`
//
//	Body string `json:”body,omitempty”`
//
//}

func DBLink (c *gin.Context) (*mongo.Client, context.Context) {

	//db 연결
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://KimDajeong:1234@cluster0-shard-00-00.fqmrn.mongodb.net:27017,cluster0-shard-00-01.fqmrn.mongodb.net:27017,cluster0-shard-00-02.fqmrn.mongodb.net:27017/website?ssl=true&replicaSet=atlas-3zdz2b-shard-0&authSource=admin&retryWrites=true&w=majority"))
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(c.Writer, "Error[00001] : 데이터베이스 연결에 실패했습니다.\n", err)
		return nil, nil
	}

	//context 작성
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(c.Writer, "Error[00002] : 데이터베이스 접근에 실패했습니다.\n", err)
		return nil, nil
	}

	return client, ctx
}

func (h *Handler) CRUD(c *gin.Context) {

	client, ctx := DBLink(c)

	//db 연결 끊기
	defer client.Disconnect(ctx)

	//crud := new(CRUD)
	//err := json.NewDecoder(c.Request.Body).Decode(crud)
	//if err != nil {
	//	c.Writer.WriteHeader(http.StatusBadRequest)
	//	fmt.Fprint(c.Writer, "Error[00003] : json 데이터 디코딩에 실피했습니다.\n", err)
	//	return
	//}
	//crud.created_at = time.Now()
	//
	//data, _ := json.Marshal(crud)
	//c.Writer.Header().Set("content-type", "application/json")
	//c.Writer.WriteHeader(http.StatusCreated)
	//fmt.Fprint(c.Writer, string(data))
	//
	////기존 특정 db 및 해당 테이블 선언
	//websitestartdatabase := client.Database(string(crud.dbname))
	//staffsCollection := websitestartdatabase.Collection(string(crud.tbname))
	//
	////데이터 입력 방식 선언 및 입력 - 하나의 데이터만 추가
	//staffResult, err := staffsCollection.InsertOne(ctx, []interface{}{crud.create})
	//fmt.Println(staffResult.InsertedID)
	//
	//fmt.Fprint(c.Writer, "Created data!")

	objA := formA{}
	objB := formB{}
	// This c.ShouldBind consumes c.Request.Body and it cannot be reused.
	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// Always an error is occurred by this because c.Request.Body is EOF now.
	} else if errB := c.ShouldBind(&objB); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	}



	fmt.Fprintf(c.Writer, objA.Foo)


	//// 몽고DB 사이트에서 찾은 Golang에서 몽고DB 추가 방법
	//post := Post{title, body}
	//
	//collection := client.Database(“my_database”).Collection(“posts”)
	//
	//insertResult, err := collection.InsertOne(context.TODO(), post)
	//
	//
	//
	//if err != nil {
	//
	//	log.Fatal(err)
	//
	//}
	//
	//fmt.Println(“Inserted post with ID:”, insertResult.InsertedID)

}

func (h *Handler) Create(c *gin.Context) {

	client, ctx := DBLink(c)

	//db 연결 끊기
	defer client.Disconnect(ctx)

	//기존 특정 db 및 해당 테이블 선언
	websitestartdatabase := client.Database("website")
	staffsCollection := websitestartdatabase.Collection("staff")
	//testsCollection := websitestartdatabase.Collection("test")

	//기존 특정 테이블 할당 1
	//cursor, err := staffsCollection.Find(ctx, bson.M{})
	//if err != nil {
	//	log.Fatal(err)
	//}

	//기존 특정 테이블 데이터 대입
	//var Staffs []bson.M
	//if err = cursor.All(ctx, &Staffs); err != nil {
	//	log.Fatal(err)
	//}

	//전체 데이터 조회
	//log.Print(Staffs)

	//데이터 입력 방식 선언 및 입력 - 하나의 데이터만 추가
	staffResult, err := staffsCollection.InsertOne(ctx, bson.D{
		{"staff_num", "21"},
		{"name", "김명철"},
		{"department", "프로그래밍팀"},
		{"title", "선임연구원"},
		{"professional_fields", bson.A{"AWS", "GO", "Backend"}},
		{"extension_num", "02-2222-2222"},
		{"phone", "010-2222-2222"},
		{"email", "2222@tritech.co.kr"},
		{"img", "https://src/tritech/members/kim/mungchual"},
		{"board", bson.D{{"open_board", bson.A{"120", "442", "992"}}}},
		{"grade", "admin"},
	})
	if err != nil {
		fmt.Fprintf(c.Writer, "Error[00003] : staff 테이블의 데이터 생성에 실패했습니다.")
	}
	fmt.Println(staffResult.InsertedID)

	//데이터 입력 방식 선언 및 입력 - 여러 데이터 추가
	//testResult, err := testsCollection.InsertMany(ctx, []interface{}{
	//	bson.D{
	//		{"member", staffResult.InsertedID},
	//		{"title", "Episode #1"},
	//		{"description", "this is the first episode"},
	//		{"duration", "25"},
	//	},
	//	bson.D{
	//		{"member", staffResult.InsertedID},
	//		{"title", "Episode #2"},
	//		{"description", "this is the second episode"},
	//		{"duration", "32"},
	//	},
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(testResult.InsertedIDs)

	//단일 인스턴스 제거
	//result, err := testsCollection.DeleteOne(ctx, bson.M{"duration": "25"})
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

	//for _, member := range Members {
	//	log.Println("Sorted :", member["id"])
	//}

	//id 정의
	//id, _ := primitive.ObjectIDFromHex("5f560e08638412174cc3c416")

	//파라미터를 통한 property update
	//result, err := testsCollection.UpdateOne(
	//	ctx,
	//	bson.M{"_id": id},
	//	bson.D{
	//		{"$set", bson.D{{"author", "Gothic"}}},
	//	},
	//)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("UpdateOne %v document(s)\n", result.ModifiedCount)

	//DB에 해당되는 스트링값을 통한 속성값 수정
	//result, err := testsCollection.UpdateMany(
	//	ctx,
	//	bson.M{"title": "Episode #2"},
	//	bson.D{
	//		{"$set", bson.D{{"author", "gothic"}}},
	//	},
	//)
	//fmt.Printf("UpdateOne %v document(s)\n", result.ModifiedCount)

	//DB에 해당되는 값 대체
	//result, err := testsCollection.ReplaceOne(
	//	ctx,
	//	bson.M{"author": "gothic"},
	//	bson.M{
	//		"title": "The Nic Raboy Story",
	//		"author": nil,
	//	},
	//)
	//fmt.Printf("UpdateOne %v document(s)\n", result.ModifiedCount)


//var u User
	//c.BindJSON(&u)
	//log.Println(u.Id)
	//
	//b := string(u.Id)
	//
	//if b == Members[0]["id"] {
	//	log.Print("Success")
	//	c.JSON(http.StatusOK, gin.H{"ID" : u.Id, "PW" : u.Pw})
	//} else {
	//	log.Print("failed!")
	//	fmt.Fprintln(c.Writer, "It's failed")
	//}

	fmt.Fprint(c.Writer, "Created data!")
}

func (h *Handler) Read(c *gin.Context) {

	client, ctx := DBLink(c)

	//db 연결 끊기
	defer client.Disconnect(ctx)

	//기존 특정 db 및 해당 테이블 선언
	websitestartdatabase := client.Database("website")
	staffsCollection := websitestartdatabase.Collection("staff")

	//기존 특정 테이블 할당 1
	cursor, err := staffsCollection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Fprintf(c.Writer, "Error[00004] : staff 테이블 정의에 실패했습니다.")
	}

	//기존 특정 테이블 데이터 대입
	var Staffs []bson.M
	if err = cursor.All(ctx, &Staffs); err != nil {
		fmt.Fprintf(c.Writer, "Error[00005] : staff 테이블 데이터 할당에 실패했습니다.")
	}

	//전체 데이터 조회
	log.Print(Staffs)

	fmt.Fprint(c.Writer, Staffs)

}

func (h *Handler) Update(c *gin.Context) {

	client, ctx := DBLink(c)

	//db 연결 끊기
	defer client.Disconnect(ctx)

	//기존 특정 db 및 해당 테이블 선언
	websitestartdatabase := client.Database("website")
	staffsCollection := websitestartdatabase.Collection("staff")

	//DB에 해당되는 스트링값을 통한 속성값 수정
	result, err := staffsCollection.UpdateMany(
		ctx,
		bson.M{"name": "김명철"},
		bson.D{
			{"$set", bson.D{{"department", "기획팀"}}},
		},
	)

	if err != nil {
		fmt.Fprintf(c.Writer, "Error[00006] : staff 테이블 데이터 수정에 실패했습니다.")
	}

	fmt.Printf("UpdateOne %v document(s)\n", result.ModifiedCount)

	fmt.Fprintf(c.Writer, "UpdateOne %v document(s)\n", result.ModifiedCount)
}

func (h *Handler) Delete(c *gin.Context) {

	client, ctx := DBLink(c)

	//db 연결 끊기
	defer client.Disconnect(ctx)

	//기존 특정 db 및 해당 테이블 선언
	websitestartdatabase := client.Database("website")
	staffsCollection := websitestartdatabase.Collection("staff")

	//단일 인스턴스 제거
	result, err := staffsCollection.DeleteOne(ctx, bson.M{"name": "김명철"})
	if err != nil {
		fmt.Fprintf(c.Writer, "Error[00007] : staff 테이블 데이터 제거에 실패했습니다.")
	}
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)

	fmt.Fprintf(c.Writer, "DeleteOne removed %v document(s)\n", result.DeletedCount)

}

