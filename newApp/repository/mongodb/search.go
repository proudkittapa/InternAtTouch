package mongodb

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"context"
	domain "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
)

func (repo *Repository)SearchDefault(ctx context.Context,field string, keyword string) []domain. {
	fmt.Println("Searching", keyword)
	var result []SuperheroQ
	cursor, err := repo.Coll.Find(ctx, bson.M{field: primitive.Regex{Pattern: keyword, Options: "i"}})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
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

func (repo *Repository)SearchByBothName(ctx context.Context,keyword string) []SuperheroQ {
	fmt.Println("Searching", keyword)
	var result []SuperheroQ
	cursor, err := repo.Coll.Find(ctx,
		bson.M{
			"$or": bson.A{
				bson.M{"name": primitive.Regex{Pattern: keyword, Options: "i"}},
				bson.M{"actual_name": primitive.Regex{Pattern: keyword, Options: "i"}},
			}})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
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

