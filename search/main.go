package main

import (
	"fmt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
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

func (this Product) ToString() string {
	result := fmt.Sprintf("id: %s", this.Id.Hex())
	result = result + fmt.Sprintf("\nname: %s", this.Name)
	result = result + fmt.Sprintf("\nprice: %0.1f", this.Price)
	result = result + fmt.Sprintf("\nquantity: %d", this.Quantity)
	result = result + fmt.Sprintf("\nstatus: %t", this.Status)
	result = result + fmt.Sprintf("\ndate: %s", this.Date.Format("2006-01-02"))
	result = result + fmt.Sprintf("\ncategory id: %s", this.CategoryId.Hex())
	result = result + fmt.Sprintf("\ncolors: %s", strings.Join(this.Colors, ", "))
	return result
}

type Product struct {
	Id         bson.ObjectId `bson:"_id"`
	Name       string        `bson:"name"`
	Price      float64       `bson:"price"`
	Quantity   int64         `bson:"quantity"`
	Status     bool          `bson:"status"`
	Date       time.Time     `bson:"date"`
	CategoryId bson.ObjectId `bson:"categoryId"`
	Colors     []string      `bson:"colors"`
}

func GetMongoDB() (*mgo.Collection, error) {
	host := "mongodb://localhost:27017"
	dbName := "superheroes"
	collection := "list"
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}
	db := session.DB(dbName).C(collection)
	//var ans superhero
	////var result_struct superhero
	//temp := db.Find(bson.M{"Name": "Superman"}).One(&ans)
	//fmt.Println(temp)
	////bsonBytes, _ := bson.Marshal(ans)
	////bson.Unmarshal(bsonBytes, &result_struct)
	//fmt.Println(ans.ToString())
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

	fmt.Println("Find the product with the name starting with lap")
	results, err := search.FindNameStartsWith("S")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, temp := range results {
			fmt.Println(temp.ToString())
			fmt.Println("---------------------------")
		}
	}
}
