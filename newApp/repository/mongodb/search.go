package mongodb

import (
	"context"
	"fmt"
	domain "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func addToArray(cursor *mongo.Cursor,err error,ctx context.Context) ([]domain.InsertQ, error) {
	var result []domain.InsertQ
	for cursor.Next(ctx) {
		var resultBson bson.M
		var resultStruct domain.InsertQ
		if err = cursor.Decode(&resultBson); err != nil {
			return result,err
		}
		bsonBytes, _ := bson.Marshal(resultBson)
		err1 := bson.Unmarshal(bsonBytes, &resultStruct)
		if err1 != nil{
			return result,err
		}
		fmt.Println(resultStruct)
		result = append(result, resultStruct)
	}
	return result,err
}

func (repo *Repository)Search(ctx context.Context,search *domain.SearchValue) (result []domain.InsertQ,err error){
	fmt.Println("Searching for ",search.Value,"in",search.Type)
	switch search.Type{
	case "name", "actual_name", "gender", "birthday", "super_power", "height", "alive":
	cursor, err := repo.Coll.Find(ctx, bson.M{search.Type: primitive.Regex{Pattern: search.Value, Options: "i"}})
		if err != nil {
			return addToArray(cursor,err,ctx)
		}
		return addToArray(cursor,err,ctx)
	case "both_name":
		cursor, err := repo.Coll.Find(ctx,
			bson.M{
				"$or": bson.A{
					bson.M{"name": primitive.Regex{Pattern: search.Value, Options: "i"}},
					bson.M{"actual_name": primitive.Regex{Pattern: search.Value, Options: "i"}},
				}})
		if err != nil {
			return addToArray(cursor,err,ctx)
		}
		return addToArray(cursor,err,ctx)
	}
	return addToArray(cursor,err,ctx)
}

