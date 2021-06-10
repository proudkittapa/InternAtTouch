package Database

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	//"strings"

	// "touch/Database"
)

type search interface{
	name(keyword string) []SuperheroQ
	actualName(keyword string) []SuperheroQ
	bothName(keyword string) []SuperheroQ
	gender(keyword string) []SuperheroQ
	superPower(keyword string) []SuperheroQ
}

type SearchValue struct {
	Value string `bson:"value"`
}

func AllSearch(feature search,keyword string) []SuperheroQ{
	return feature.name(keyword)
}

func (heroDB *dbconfig)name(keyword string) []SuperheroQ {
	fmt.Println("Searching", keyword)
	var result []SuperheroQ
	cursor, err := heroDB.Coll.Find(heroDB.Ctx, bson.M{"name": primitive.Regex{Pattern: keyword, Options: "i"}})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(heroDB.Ctx) {
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

func (heroDB *dbconfig)actualName(keyword string) []SuperheroQ {
	fmt.Println("Searching", keyword)
	var result []SuperheroQ
	cursor, err := heroDB.Coll.Find(heroDB.Ctx, bson.M{"actual_name": primitive.Regex{Pattern: keyword, Options: "i"}})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(heroDB.Ctx) {
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

func (heroDB *dbconfig)bothName(keyword string) []SuperheroQ {
	fmt.Println("Searching", keyword)
	var result []SuperheroQ
	cursor, err := heroDB.Coll.Find(heroDB.Ctx,
		bson.M{
			"$or": bson.A{
				bson.M{"name": primitive.Regex{Pattern: keyword, Options: "i"}},
				bson.M{"actual_name": primitive.Regex{Pattern: keyword, Options: "i"}},
			}})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(db.Ctx) {
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

func (heroDB *dbconfig)gender(keyword string) []SuperheroQ {
	fmt.Println("Searching", keyword)
	var result []SuperheroQ
	cursor, err := heroDB.Coll.Find(heroDB.Ctx, bson.M{"gender": primitive.Regex{Pattern: keyword, Options: "i"}})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(heroDB.Ctx) {
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

func (heroDB *dbconfig)superPower(keyword string) []SuperheroQ {
	fmt.Println("Searching", keyword)
	var result []SuperheroQ
	cursor, err := heroDB.Coll.Find(heroDB.Ctx, bson.M{"super_power": primitive.Regex{Pattern: keyword, Options: "i"}})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(heroDB.Ctx) {
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

