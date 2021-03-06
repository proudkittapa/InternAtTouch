
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Person struct {
	Name string
}

var name string

func main(){

	http.HandleFunc("/", test)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

func test(rw http.ResponseWriter, req *http.Request ){
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var t Person
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	log.Println(t.Name)
	name = t.Name
	db()
	//collection = client.Database("test").Collection("your_collection_name")
	//fmt.Println(collection)
}

func db(){
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://@localhost:28017"))
	collection := client.Database("test").Collection("pond")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	fmt.Println("name", name)
	p:=Person{name}
	insertResult, err := collection.InsertOne(context.TODO(), p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

}