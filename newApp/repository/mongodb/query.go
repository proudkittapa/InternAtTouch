package mongodb

import (
	"context"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func (repo *Repository)Create(ctx context.Context, figure interface{}) (  err error){
	//initID := goxid.New()
	//figure.ID  = initID.Gen()
<<<<<<< HEAD
	_, err = repo.Coll.InsertOne(ctx, figure)
=======
	, err = repo.Coll.InsertOne(ctx, figure)
>>>>>>> a8d1e4e56b3781069d1476e2788c7c6e5f7d8d62
	return  err
}

func (repo *Repository)Delete(ctx context.Context, id interface{}) (err error){
<<<<<<< HEAD
	_, err = repo.Coll.DeleteOne(ctx, bson.M{"_id": id})
=======
	, err = repo.Coll.DeleteOne(ctx, bson.M{"id": id})
>>>>>>> a8d1e4e56b3781069d1476e2788c7c6e5f7d8d62
	return err
}

func (repo *Repository)Update(ctx context.Context, figure interface{}, id string) (err error){
<<<<<<< HEAD
	_, err = repo.Coll.UpdateOne(ctx, bson.M{"_id": id}, bson.D{{"$set", figure},},)
=======
	, err = repo.Coll.UpdateOne(ctx, bson.M{"_id": id}, bson.D{{"$set", figure},},)
>>>>>>> a8d1e4e56b3781069d1476e2788c7c6e5f7d8d62
	return err
}

func (repo *Repository)View(ctx context.Context, id string) (domain.InsertQ, error){
	var resultBson bson.D
	var resultStruct domain.InsertQ
<<<<<<< HEAD
	err := repo.Coll.FindOne(ctx, bson.D{{"_id", id}}).Decode(&resultBson)
	bsonBytes, _ := bson.Marshal(resultBson)
=======
	err := repo.Coll.FindOne(ctx, bson.D{{"id", id}}).Decode(&resultBson)
	bsonBytes,  := bson.Marshal(resultBson)
>>>>>>> a8d1e4e56b3781069d1476e2788c7c6e5f7d8d62
	bson.Unmarshal(bsonBytes, &resultStruct)
	return resultStruct ,err
}

func (repo *Repository)ViewAll(ctx context.Context, perPage int, page int)([]domain.InsertQ, error) {
	skip := int64(page * perPage)
	limit := int64(perPage)
	opts := options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}
	cursor, err := repo.Coll.Find(nil, bson.M{}, &opts)
	var display []domain.InsertQ
	for cursor.Next(ctx) {
		var resultBson bson.D
		var resultStruct domain.InsertQ
		if err = cursor.Decode(&resultBson); err != nil {
			log.Fatal(err)
		}
		bsonBytes, _ := bson.Marshal(resultBson)
		bson.Unmarshal(bsonBytes, &resultStruct)
		display = append(display, resultStruct)
	}
	return display, err
}