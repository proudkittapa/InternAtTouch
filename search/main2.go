package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type superhero struct {
	ID 			int		`bson:"ID"`
	Name  		string	`bson:"Name"`
	Actual string	`bson:"Actual_name"`
	Gender 		string	`bson:"Gender"`
	Age 		int		`bson:"Age"`
	Superpower string	`bson:"Super_power"`
}

func (this superhero) ToString() string {
	result := fmt.Sprintf("ID: %d", this.ID)
	result = result + fmt.Sprintf("\nName: %s", this.Name)
	result = result + fmt.Sprintf("\nActual name: %s", this.Actual)
	result = result + fmt.Sprintf("\nGender: %s", this.Gender)
	result = result + fmt.Sprintf("\nAge: %d", this.Age)
	result = result + fmt.Sprintf("\ndate: %s", this.Superpower)
	return result
}

func GetMongoDB() (Database.Coll, error) {
	host := "mongodb://localhost:27017"
	dbName := "superheroes"
	collection := "list"
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}
	db := session.DB(dbName).C(collection)
	return db, nil
}

type hero struct {
}

func (this hero) FindNameStartsWith(keyword string) ([]superhero, error) {
	db, err := GetMongoDB()
	var heroes []superhero
	if err != nil {
		return heroes, err
	} else {
		err2 := db.Find(bson.M{
			"Name": bson.RegEx{
				Pattern: "^" + keyword,
				Options: "i",
			},
		}).All(&heroes)
		fmt.Println(heroes)
		if err2 != nil {
			return heroes, err
		} else {
			return heroes, nil
		}
	}
}

func main(){
	var search hero
	results, err := search.FindNameStartsWith("z")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, temp := range results {
			fmt.Println(temp.ToString())
			fmt.Println("---------------------------")
		}
	}
}
