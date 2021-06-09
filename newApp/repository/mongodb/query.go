package mongodb

import (
	"context"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *Repository)Create(ctx context.Context, figure interface{}) (  err error){
	_, err = repo.Coll.InsertOne(ctx, figure)
	return  err
}

func (repo *Repository)Delete(ctx context.Context, id interface{}) (err error){
	_, err = repo.Coll.DeleteOne(ctx, bson.M{"id": id})
	return err
}

func (repo *Repository)Update(ctx context.Context, figure interface{}, id string) (err error){
	_, err = repo.Coll.UpdateOne(ctx, bson.M{"_id": id}, bson.D{{"$set", figure},},)
	return err
}

func (repo *Repository)View(ctx context.Context, id string) (domain.InsertQ, error){
	var resultBson bson.D
	var resultStruct domain.InsertQ
	err := repo.Coll.FindOne(ctx, bson.D{{"id", id}}).Decode(&resultBson)
	bsonBytes,_ := bson.Marshal(resultBson)
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
	return AddToArray(cursor, err, ctx)
}