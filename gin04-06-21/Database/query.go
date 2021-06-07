package Database

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type SuperheroQ struct {
	ID          int    		`bson:"ID"`
	Name        string 		`bson:"Name"`
	ActualName 	string 		`bson:"ActualName"`
	Gender      string 		`bson:"Gender"`
	BirthDate	string		`bson:"BirthDate"`
	Height      int    		`bson:"Height"`
	SuperPower  []string 	`bson:"SuperPower"`
	Alive		bool		`bson:"Alive"`
}

func MaxId() int {
	var result bson.M
	var currID SuperheroQ
	opts := options.FindOne().SetSort(bson.D{{"ID", -1}})
	err := Coll.FindOne(Ctx, bson.D{{}}, opts).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	bsonBytes, _ := bson.Marshal(result)
	bson.Unmarshal(bsonBytes, &currID)
	return currID.ID
}

func udstr(id int, key string, valStr string) {
	_, err := Coll.UpdateOne(
		Ctx,
		bson.M{"ID": id},
		bson.D{
			{"$set", bson.D{{key, valStr}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func CheckExistID(id int) bool {
	count, err := Coll.CountDocuments(Ctx, bson.D{{"ID", id}})
	if err != nil {
		panic(err)
	}
	if count >= 1 {
		return true
	}
	return false
}

func CheckExistName(name string) bool {
	count, err := Coll.CountDocuments(Ctx, bson.D{{"Name", name}})
	if err != nil {
		panic(err)
	}
	if count >= 1 {
		return true
	}
	return false
}

func Insert(figure SuperheroQ) {
	bd , err := time.Parse("2006-01-02", figure.BirthDate)
	_, err = Coll.InsertOne(Ctx, bson.D{
		{"$inc", bson.D{{"ID", 1}}},
		{"Name", figure.Name},
		{"Actual_name", figure.ActualName},
		{"Gender", figure.Gender},
		{"BirthDate", bd},
		{"Height", figure.Height},
		{"Super_power", figure.SuperPower},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func Delete(id int) {
	_, err := Coll.DeleteOne(Ctx, bson.M{"ID": id})
	if err != nil {
		log.Fatal(err)
	}
}

func Update(figure SuperheroQ, id int) {
	//id := figure.ID
	if figure.Name != "" {
		udstr(id, "Name", figure.Name)
	}
	if figure.ActualName != "" {
		udstr(id, "Actual_name", figure.ActualName)
	}
	if figure.Gender != "" {
		udstr(id, "Gender",  figure.Gender)
	}
	if figure.Height != -1 {
		_, err := Coll.UpdateOne(
			Ctx,
			bson.M{"ID": id},
			bson.D{
				{"$set", bson.D{{"Height", figure.Height}}},
			},
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	if figure.SuperPower != nil {
		_, err := Coll.UpdateOne(
			Ctx,
			bson.M{"ID": id},
			bson.D{
				{"$set", bson.D{{"SuperPower", figure.SuperPower}}},
			},
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	if figure.BirthDate != "" {
		bd , err := time.Parse("2006-01-02", figure.BirthDate)
		_, err = Coll.UpdateOne(
			Ctx,
			bson.M{"ID": id},
			bson.D{
				{"$set", bson.D{{"BirthDate",bd }}},
			},
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func View(id int) SuperheroQ {
	var resultBson bson.M
	var resultStruct SuperheroQ
	err := Coll.FindOne(Ctx, bson.D{{"ID", id}}).Decode(&resultBson)
	if err != nil {
		log.Fatal(err)
	}
	bsonBytes, _ := bson.Marshal(resultBson)
	bson.Unmarshal(bsonBytes, &resultStruct)
	fmt.Println(resultStruct)
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
		var resultBson bson.M
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
			var resultBson bson.M
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
