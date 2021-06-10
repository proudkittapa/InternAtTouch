package mongodb

import (
	"context"
<<<<<<< Updated upstream
	"log"

=======
	"errors"
>>>>>>> Stashed changes
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *Repository) Create(ctx context.Context, figure interface{}) (err error) {
	_, err = repo.Coll.InsertOne(ctx, figure)
	return err
}

<<<<<<< Updated upstream
func (repo *Repository) Delete(ctx context.Context, id string) (err error) {
=======
func (repo *Repository)Delete(ctx context.Context, id string) (err error){
	_, err = repo.checkExistID(ctx, id)
	if err != nil {
		return err
	}
>>>>>>> Stashed changes
	_, err = repo.Coll.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (repo *Repository) Update(ctx context.Context, figure interface{}, id string) (err error) {
	_, err = repo.Coll.UpdateOne(ctx, bson.M{"_id": id}, bson.D{{"$set", figure}})
	return err
}

<<<<<<< Updated upstream
func (repo *Repository) View(ctx context.Context, id string) (domain.InsertQ, error) {
	var resultBson bson.D
	var resultStruct domain.InsertQ
	err := repo.Coll.FindOne(ctx, bson.D{{"_id", id}}).Decode(&resultBson)
	bsonBytes, _ := bson.Marshal(resultBson)
=======
func (repo *Repository)View(ctx context.Context, id string) (resultStruct domain.InsertQ, err error){
	_, err = repo.checkExistID(ctx, id)
	if err != nil {
		return resultStruct, err
	}
	var resultBson bson.D
	err = repo.Coll.FindOne(ctx, bson.D{{"_id", id}}).Decode(&resultBson)
	bsonBytes,_ := bson.Marshal(resultBson)
>>>>>>> Stashed changes
	bson.Unmarshal(bsonBytes, &resultStruct)
	return resultStruct, err
}

func (repo *Repository) ViewAll(ctx context.Context, perPage int, page int) ([]domain.InsertQ, error) {
	skip := int64(page * perPage)
	limit := int64(perPage)
	opts := options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}
	cursor, err := repo.Coll.Find(nil, bson.M{}, &opts)
	return AddToArray(cursor, err, ctx)
}

<<<<<<< Updated upstream
func (repo *Repository) CheckExistID(ctx context.Context, ID string) (bool, error) {
	count, err := repo.Coll.CountDocuments(ctx, bson.D{{"_id", ID}})
=======
func (repo *Repository) checkExistID(ctx context.Context, id string) (bool, error){
	count, err := repo.Coll.CountDocuments(ctx, bson.D{{"_id", id}})
>>>>>>> Stashed changes
	if count < 1 {
		err = errors.New("ID does not exist")
		return false, err
	}
	return true, err
}

func (repo *Repository) CheckExistName(ctx context.Context, name string) (bool, error) {
	log.Println("checkexistname")
	count, err := repo.Coll.CountDocuments(ctx, bson.D{{"name", name}})
	if count < 1 {
		return false, err
	}
	return true, err
}

func (repo *Repository) CheckExistActualName(ctx context.Context, actualName string) (bool, error) {
	count, err := repo.Coll.CountDocuments(ctx, bson.D{{"actual_name", actualName}})
	if count < 1 {
		return false, err
	}
	return true, err
}
