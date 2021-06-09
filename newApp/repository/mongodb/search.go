package mongodb

import (
	"context"
	"fmt"
	domain "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func addToArray(cursor *mongo.Cursor,err error,ctx context.Context) []domain.InsertQ{
	var result []domain.InsertQ
	for cursor.Next(ctx) {
		var resultBson bson.M
		var resultStruct domain.InsertQ
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

func (repo *Repository)SearchDefault(ctx context.Context,field string, keyword string) []domain.InsertQ{
	fmt.Println("Searching for ",keyword,"in",field)
	cursor, err := repo.Coll.Find(ctx, bson.M{field: primitive.Regex{Pattern: keyword, Options: "i"}})
	if err != nil {
		log.Fatal(err)
	}
	return addToArray(cursor,err,ctx)
}

func (repo *Repository)SearchByBothName(ctx context.Context,field string,keyword string) []domain.InsertQ {
	fmt.Println("Searching for ",keyword,"in",field)
	cursor, err := repo.Coll.Find(ctx,
		bson.M{
			"$or": bson.A{
				bson.M{"name": primitive.Regex{Pattern: keyword, Options: "i"}},
				bson.M{"actual_name": primitive.Regex{Pattern: keyword, Options: "i"}},
			}})
	if err != nil {
		log.Fatal(err)
	}
	return addToArray(cursor,err,ctx)
}
