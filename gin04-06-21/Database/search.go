package Database

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"log"
	// "touch/Database"
)

type SuperheroQ struct {
	ID         string      `bson:"_id" json:"id"`
	Name       string   `bson:"name" json:"name" binding:"required"`
	ActualName string   `bson:"actual_name" json:"actual_name"`
	Gender     string   `bson:"gender" json:"gender" binding:"required"`
	BirthDate  int64    `bson:"birth_date" json:"birth_date"`
	Height     int      `bson:"height" json:"height"`
	SuperPower []string `bson:"super_power" json:"super_power"`
	Alive      bool     `bson:"alive" json:"alive"`
}

type SearchValue struct {
	Value string `bson:"value"`
}

func SearchName(keyword string) []SuperheroQ {
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

func SearchActualName(keyword string) []SuperheroQ {
	fmt.Println("Searching", keyword)
	var result []SuperheroQ
	cursor, err := Coll.Find(Ctx, bson.M{"ActualName": primitive.Regex{Pattern: "^" + keyword + ".*", Options: "i"}})
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

func SearchContainName(keyword string) []SuperheroQ {
	fmt.Println("Searching", keyword)
	var result []SuperheroQ
	cursor, err := Coll.Find(Ctx, bson.M{"Name": primitive.Regex{Pattern:keyword, Options: "i"}})
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

func SearchContainActualName(keyword string) []SuperheroQ {
	fmt.Println("Searching", keyword)
	var result []SuperheroQ
	cursor, err := Coll.Find(Ctx, bson.M{"ActualName": primitive.Regex{Pattern:keyword, Options: "i"}})
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