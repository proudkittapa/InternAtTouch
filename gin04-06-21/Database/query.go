package Database

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Superhero_q struct {
	ID 			int		`bson:"ID"`
	Name  		string	`bson:"Name"`
	Actual_name string	`bson:"Actual_name"`
	Gender 		string	`bson:"Gender"`
	Age 		int		`bson:"Age"`
	Super_power string	`bson:"Super_power"`
}


func max_id() int {
	var result bson.M
	var curr_id Superhero_q
	opts := options.FindOne().SetSort(bson.D{{"ID", -1}})
	err := Coll.FindOne(Ctx, bson.D{{}}, opts).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	bsonBytes, _ := bson.Marshal(result)
	bson.Unmarshal(bsonBytes, &curr_id)
	return curr_id.ID
}

func ud_one(id int, key string, val_int int, val_str string){
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

func Check_exist_ID(id int) bool{
	count, err := Coll.CountDocuments(Ctx, bson.D{{"ID", id}})
	if err != nil {
		panic(err)
	}
	if count >= 1 {
		return true
	}
	return false
}

func Check_exist_Name(name string) bool{
	count, err := Coll.CountDocuments(Ctx, bson.D{{"Name", name}})
	if err != nil {
		panic(err)
	}
	if count >= 1 {
		return true
	}
	return false
}


func Insert(figure Superhero_q){
	_, err := Coll.InsertOne(Ctx, bson.D{
		{"ID", max_id()+1},
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

func Delete(id int){
	_, err := Coll.DeleteOne(Ctx, bson.M{"ID": id})
	if err != nil {
		log.Fatal(err)
	}
}

func Update(figure Superhero_q, id int){
	//id := figure.ID
	if figure.Name != ""{
		ud_one(id, "Name", -1, figure.Name)
	}
	if figure.Actual_name != ""{
		ud_one(id, "Actual_name", -1, figure.Actual_name)
	}
	if figure.Gender != ""{
		ud_one(id, "Gender", -1, figure.Gender)
	}
	if figure.Age != -1{
		ud_one(id, "Age",  figure.Age, "")
	}
	if figure.Super_power != ""{
		ud_one(id, "Super_power", -1, figure.Super_power)
	}
}

func View(id int) Superhero_q{
	var result_bson bson.M
	var result_struct Superhero_q
	err := Coll.FindOne(Ctx, bson.D{{"ID", id}}).Decode(&result_bson)
	if err != nil {
		log.Fatal(err)
	}
	bsonBytes, _ := bson.Marshal(result_bson)
	bson.Unmarshal(bsonBytes, &result_struct)
	fmt.Println(result_struct)
	return result_struct
}

func Viewall(limit int, offset int) []Superhero_q{
	var display []Superhero_q
	cursor, err := Coll.Find(Ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(Ctx)
	count := 1
	start := (offset)*limit
	stop := (offset+1)*limit
	for cursor.Next(Ctx) {
		if count > start && count <= stop{
			var result_bson bson.M
			var result_struct Superhero_q
			if err = cursor.Decode(&result_bson); err != nil {
				log.Fatal(err)
			}
			bsonBytes, _ := bson.Marshal(result_bson)
			bson.Unmarshal(bsonBytes, &result_struct)
			display = append(display, result_struct)

			if count == stop{
				fmt.Println(display)
				return display
			}
		}
		count += 1
	}
	fmt.Println(display)
	return display
}


