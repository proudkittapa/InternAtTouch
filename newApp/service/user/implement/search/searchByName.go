package search

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func (keyword searchValue)name(keyword string) []SuperheroQ {
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
