package search

import (
	"fmt"
	"touch/Database"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (this Database.Superhero_q) ToString() string {
	result := fmt.Sprintf("ID: %d", this.ID)
	result = result + fmt.Sprintf("\nName: %s", this.Name)
	result = result + fmt.Sprintf("\nActual name: %s", this.Actual)
	result = result + fmt.Sprintf("\nGender: %s", this.Gender)
	result = result + fmt.Sprintf("\nAge: %d", this.Age)
	result = result + fmt.Sprintf("\ndate: %s", this.Superpower)
	return result
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
	return db, nil
}

type hero struct {
}

func (this hero) FindNameStartsWith(keyword string) ([]Database.Superhero_q, error) {
	db, err := GetMongoDB()
	var heroes []Database.Superhero_q
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

// func main() {
// 	var search hero
// 	results, err := search.FindNameStartsWith("z")
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		for _, temp := range results {
// 			fmt.Println(temp.ToString())
// 			fmt.Println("---------------------------")
// 		}
// 	}
// }
