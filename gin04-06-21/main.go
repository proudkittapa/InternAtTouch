package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"touch/Database"
)

type Person struct {
	Name string
}

type Sp struct {
	ID 			int		`bson:"ID"`
	Name  		string	`bson:"Name"`
	Actual_name string	`bson:"Actual_name"`
	Gender 		string	`bson:"Gender"`
	Age 		int		`bson:"Age"`
	Super_power string	`bson:"Super_power"`
}



func main(){
	Database.InitDB()
	http.HandleFunc("/", test)
	log.Fatal(http.ListenAndServe(":8080", nil))

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
	db(t)
}

func db(name Person) {
	fmt.Println("name", name.Name)
	insertResult, err := Database.Coll.InsertOne(context.TODO(), name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}