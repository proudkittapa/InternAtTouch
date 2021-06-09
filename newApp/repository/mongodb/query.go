package mongodb

import (
	"context"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
	goxid "github.com/touchtechnologies-product/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func (repo *Repository)Create(ctx context.Context, figure domain.InsertQ) (ID string,  err error){
	initID := goxid.New()
	figure.ID  = initID.Gen()

	_, err = repo.Coll.InsertOne(ctx, figure)
	if err != nil {
		log.Fatal("err3: ", err)
	}
	return figure.ID , err
}

func (repo *Repository)Delete(ctx context.Context, id int) (err error){
	_, err = repo.Coll.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (repo *Repository)Update(ctx context.Context, figure domain.InsertQ) (err error){
	_, err = repo.Coll.UpdateOne(ctx, bson.M{"_id": figure.ID}, bson.D{{"$set", figure},},)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (repo *Repository)View(ctx context.Context, id int) (domain.InsertQ, error){
	var resultBson bson.D
	var resultStruct domain.InsertQ
	err := repo.Coll.FindOne(ctx, bson.D{{"_id", id}}).Decode(&resultBson)
	if err != nil {
		log.Fatal(err)
	}
	bsonBytes, _ := bson.Marshal(resultBson)
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
