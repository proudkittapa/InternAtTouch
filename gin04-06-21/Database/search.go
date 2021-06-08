package Database

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"log"
	// "touch/Database"
)

type SuperheroQ struct {
	ID         string   `bson:"_id"`
	Name       string   `bson:"Name"`
	ActualName string   `bson:"ActualName"`
	Gender     string   `bson:"Gender"`
	BirthDate  string   `bson:"BirthDate"`
	Height     int      `bson:"Height"`
	SuperPower []string `bson:"SuperPower"`
	Alive      bool     `bson:"Alive"`
}

type SearchValue struct {
	Value string `bson:"value"`
}

func Search(keyword string) []SuperheroQ {
	fmt.Println("Searching", keyword)
	var result []SuperheroQ
	cursor, err := Coll.Find(Ctx, bson.M{"Name": primitive.Regex{Pattern: "^" + keyword + ".*", Options: "i"}})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(Ctx) {
		var resultBson bson.M
		var resultStruct SuperheroQ
		if err = cursor.Decode(&resultBson); err != nil {
			log.Fatal(err)
		}
		bsonBytes, _ := bson.Marshal(resultBson)
		bson.Unmarshal(bsonBytes, &resultStruct)
		fmt.Println(resultStruct)
		result = append(result, resultStruct)
	}
	return result
}
