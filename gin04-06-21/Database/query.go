package Database

import (
	"fmt"
	"log"

	goxid "github.com/touchtechnologies-product/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//func MaxId() string {
//	var result bson.M
//	var currID SuperheroQ
//	opts := options.FindOne().SetSort(bson.D{{"_id", -1}})
//	err := Coll.FindOne(Ctx, bson.D{{}}, opts).Decode(&result)
//	if err != nil {
//		log.Fatal(err)
//	}
//	bsonBytes, _ := bson.Marshal(result)
//	bson.Unmarshal(bsonBytes, &currID)
//	return currID.ID
//}

func udstr(id string, key string, valStr string) {
	_, err := Coll.UpdateOne(
		Ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{key, valStr}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func CheckExistID(id string) bool {
	count, err := Coll.CountDocuments(Ctx, bson.D{{"_id", id}})
	if err != nil {
		log.Fatal("err1: ", err)
		return false
	}
	if count >= 1 {
		return false
	}
	return true
}

func CheckExistName(name string) bool {
	count, err := Coll.CountDocuments(Ctx, bson.D{{"name", name}})
	if err != nil {
		log.Fatal("err2: ", err)
		return false
	}
	if count >= 1 {
		return false
	}
	return true
}

func CheckExistActualName(actualName string) bool {
	count, err := Coll.CountDocuments(Ctx, bson.D{{"actual_name", actualName}})
	if err != nil {
		log.Fatal("err2: ", err)
		return false
	}
	if count >= 1 {
		return false
	}
	return true
}

func Insert(figure SuperheroQ) {
	initID := goxid.New()
	idGen := initID.Gen()
	_, err := Coll.InsertOne(Ctx, bson.D{
		{"_id", idGen},
		{"name", figure.Name},
		{"actual_name", figure.ActualName},
		{"gender", figure.Gender},
		{"birth_date", figure.BirthDate},
		{"height", figure.Height},
		{"super_power", figure.SuperPower},
		{"alive", figure.Alive},
	})
	if err != nil {
		log.Fatal("err3: ", err)
	}
}

func Delete(id string) {
	_, err := Coll.DeleteOne(Ctx, bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}
}

func Update(figure SuperheroQ, id string) {
	if figure.Name != "" {
		udstr(id, "name", figure.Name)
	}
	if figure.ActualName != "" {
		udstr(id, "actual_name", figure.ActualName)
	}
	if figure.Gender != "" {
		udstr(id, "gender", figure.Gender)
	}
	if figure.Height != -1 {
		_, err := Coll.UpdateOne(
			Ctx,
			bson.M{"_id": id},
			bson.D{
				{"$set", bson.D{{"height", figure.Height}}},
			},
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	if figure.SuperPower != nil {
		_, err := Coll.UpdateOne(
			Ctx,
			bson.M{"_id": id},
			bson.D{
				{"$set", bson.D{{"super_power", bson.A{figure.SuperPower}}}},
			},
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	if figure.BirthDate != 0 {
		_, err := Coll.UpdateOne(
			Ctx,
			bson.M{"_id": id},
			bson.D{
				{"$set", bson.D{{"birth_date", figure.BirthDate}}},
			},
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func View1(id string) SuperheroQ {
	// TODO : fix the bug of returning the incorrect output
	var resultBson bson.D
	var resultStruct SuperheroQ
	err := Coll.FindOne(Ctx, bson.D{{"_id", id}}).Decode(&resultBson)
	if err != nil {
		log.Fatal(err)
	}
	bsonBytes, _ := bson.Marshal(resultBson)
	fmt.Println(resultBson)
	bson.Unmarshal(bsonBytes, &resultStruct)
	fmt.Println(resultStruct)
	return resultStruct
}

func View(id string) SuperheroQ {
	// TODO : fix the bug of returning the incorrect output
	var resultBson bson.D
	var resultStruct SuperheroQ
	err := Coll.FindOne(Ctx, bson.D{{"_id", id}}).Decode(&resultBson)
	if err != nil {
		log.Fatal(err)
	}
	bsonBytes, _ := bson.Marshal(resultBson)
	bson.Unmarshal(bsonBytes, &resultStruct)
	return resultStruct
}

func ViewByPage(perPage int, page int) []SuperheroQ {
	skip := int64(page * perPage)
	limit := int64(perPage)
	opts := options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}

	cursor, err := Coll.Find(nil, bson.M{}, &opts)
	var display []SuperheroQ
	for cursor.Next(Ctx) {
		var resultBson bson.D
		var resultStruct SuperheroQ
		if err = cursor.Decode(&resultBson); err != nil {
			log.Fatal(err)
		}
		bsonBytes, _ := bson.Marshal(resultBson)
		bson.Unmarshal(bsonBytes, &resultStruct)
		display = append(display, resultStruct)
	}
	return display
}

func ViewByGt(perPage int, page int) []SuperheroQ {
	// TODO make the paging by using $gt
	skip := int64(page * perPage)
	limit := int64(perPage)
	opts := options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}

	cursor, err := Coll.Find(nil, bson.M{}, &opts)
	var display []SuperheroQ
	for cursor.Next(Ctx) {
		var resultBson bson.D
		var resultStruct SuperheroQ
		if err = cursor.Decode(&resultBson); err != nil {
			log.Fatal(err)
		}
		bsonBytes, _ := bson.Marshal(resultBson)
		bson.Unmarshal(bsonBytes, &resultStruct)
		display = append(display, resultStruct)
	}
	return display
}

func ViewAll(limit int, offset int) []SuperheroQ {
	//Manual version of paging
	var display []SuperheroQ
	cursor, err := Coll.Find(Ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(Ctx)
	count := 1
	start := (offset) * limit
	stop := (offset + 1) * limit
	for cursor.Next(Ctx) {
		if count > start && count <= stop {
			var resultBson bson.D
			var resultStruct SuperheroQ
			if err = cursor.Decode(&resultBson); err != nil {
				log.Fatal(err)
			}
			bsonBytes, _ := bson.Marshal(resultBson)
			bson.Unmarshal(bsonBytes, &resultStruct)
			display = append(display, resultStruct)

			if count == stop {
				fmt.Println(display)
				return display
			}
		}
		count += 1
	}
	fmt.Println(display)
	return display
}
