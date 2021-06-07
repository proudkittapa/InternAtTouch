package Database

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"log"
	// "touch/Database"
)

type SearchValue struct {
	Value string `bson:"value"`
}

func Search(keyword string) []Superhero_q {
	fmt.Println("Searching", keyword)
	var result []Superhero_q
	cursor, err := Coll.Find(Ctx, bson.M{"Name": primitive.Regex{Pattern: "^" + keyword + ".*", Options: "i"}})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(Ctx) {
		var result_bson bson.M
		var result_struct Superhero_q
		if err = cursor.Decode(&result_bson); err != nil {
			log.Fatal(err)
		}
		bsonBytes, _ := bson.Marshal(result_bson)
		bson.Unmarshal(bsonBytes, &result_struct)
		fmt.Println(result_struct)
		result = append(result, result_struct)
	}
	return result
}
