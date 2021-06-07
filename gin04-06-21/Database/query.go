package Database

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SuperheroQ struct {
	ID          int    `bson:"ID"`
	Name        string `bson:"Name"`
	Actual_name string `bson:"Actual_name"`
	Gender      string `bson:"Gender"`
	Age         int    `bson:"Age"`
	Super_power string `bson:"Super_power"`
}

func MaxId() int {
	var result bson.M
	var curr_id SuperheroQ
	opts := options.FindOne().SetSort(bson.D{{"ID", -1}})
	err := Coll.FindOne(Ctx, bson.D{{}}, opts).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	bsonBytes, _ := bson.Marshal(result)
	bson.Unmarshal(bsonBytes, &curr_id)
	return curr_id.ID
}

func udOne(id int, key string, val_int int, val_str string) {
	if val_int == -1 {
		_, err := Coll.UpdateOne(
			Ctx,
			bson.M{"ID": id},
			bson.D{
				{"$set", bson.D{{key, val_str}}},
			},
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	if val_str == "" {
		_, err := Coll.UpdateOne(
			Ctx,
			bson.M{"ID": id},
			bson.D{
				{"$set", bson.D{{key, val_int}}},
			},
		)
		if err != nil {
			log.Fatal(err)
		}
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
	_, err := Coll.InsertOne(Ctx, bson.D{
		{"ID", MaxId() + 1},
		{"Name", figure.Name},
		{"Actual_name", figure.Actual_name},
		{"Gender", figure.Gender},
		{"Age", figure.Age},
		{"Super_power", figure.Super_power},
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
		udOne(id, "Name", -1, figure.Name)
	}
	if figure.Actual_name != "" {
		udOne(id, "Actual_name", -1, figure.Actual_name)
	}
	if figure.Gender != "" {
		udOne(id, "Gender", -1, figure.Gender)
	}
	if figure.Age != -1 {
		udOne(id, "Age", figure.Age, "")
	}
	if figure.Super_power != "" {
		udOne(id, "Super_power", -1, figure.Super_power)
	}
}

func View(id int) SuperheroQ {
	var result_bson bson.M
	var result_struct SuperheroQ
	err := Coll.FindOne(Ctx, bson.D{{"ID", id}}).Decode(&result_bson)
	if err != nil {
		log.Fatal(err)
	}
	bsonBytes, _ := bson.Marshal(result_bson)
	bson.Unmarshal(bsonBytes, &result_struct)
	fmt.Println(result_struct)
	return result_struct
}

func View_byPage(perPage int, page int) []SuperheroQ {
	skip := int64(page * perPage)
	limit := int64(perPage)
	opts := options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}

	cursor, err := Coll.Find(nil, bson.M{}, &opts)
	var display []SuperheroQ
	for cursor.Next(Ctx) {
		var result_bson bson.M
		var result_struct SuperheroQ
		if err = cursor.Decode(&result_bson); err != nil {
			log.Fatal(err)
		}
		bsonBytes, _ := bson.Marshal(result_bson)
		bson.Unmarshal(bsonBytes, &result_struct)
		display = append(display, result_struct)
	}
	return display
}

func Viewall(limit int, offset int) []SuperheroQ {
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
			var result_bson bson.M
			var result_struct SuperheroQ
			if err = cursor.Decode(&result_bson); err != nil {
				log.Fatal(err)
			}
			bsonBytes, _ := bson.Marshal(result_bson)
			bson.Unmarshal(bsonBytes, &result_struct)
			display = append(display, result_struct)

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
