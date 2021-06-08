package main

import (
	"context"
	goxid "github.com/touchtechnologies-product/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Sp struct {
	Name        string 		`bson:"Name"`
	ActualName 	string 		`bson:"ActualName"`
	Gender      string 		`bson:"Gender"`
	BirthDate	string		`bson:"BirthDate"`
	Height      int    		`bson:"Height"`
	SuperPower  []string 	`bson:"SuperPower"`
	Alive		bool		`bson:"Alive"`
}

var Sp_list =  []Sp{
	{"Spider-Man", "Peter Parker", "Male", "2001-08-10", 178, []string{"Web-shooting"}, true},
	{"Batman", "Bruce Wayne", "Male", "1978-04-17", 188, []string{"Rich"}, true},
	{"Superman", "Clark Kent", "Male", "1977-04-18", 191, []string{"Flight", "Strength"}, true},
	{"Wonder woman", "Diana Prince", "Female", "1941-03-22", 178, []string{"Agility" , "Strength"}, true},
	{"Doctor Strange", "Stephen Vincent Strange", "Male", "1930-11-18", 183, []string{"Magic"}, true},
	{"Iron man", "Tony Stark", "Male", "1970-05-29", 185, []string{"Genius", "super-suit"}, false},
	{"Black Widow", "Natasha Romanoff", "Female", "1984-11-22", 170, []string{"Expert spy"}, false},
	{"Scarlet Witch", "Wanda Maximoff", "Female", "1976-02-10", 170, []string{"Energy manipulation"}, true},
	{"Harley Quinn", "Dr. Harleen Quinzel", "Female", "1999-06-26", 168, []string{"Immunity and Strength"}, true},
	{"Captain America", "Steve Rogers", "Male", "1918-07-04", 188, []string{"Immunity", "Strength"}, true},
}

func main(){
	uri := "mongodb://touch:touchja@localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	collection := client.Database("superhero").Collection("lists")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}


	defer client.Disconnect(ctx)
	for _ ,v := range Sp_list{
		birthdate , err := time.Parse("2006-01-02", v.BirthDate)
		if err != nil {
			log.Fatal(err)
		}
		initID := goxid.New()
		idGen := initID.Gen()
		_, err = collection.InsertOne(ctx, bson.D{
			{"_id", idGen},
			{"Name", v.Name},
			{"ActualName", v.ActualName},
			{"Gender", v.Gender},
			{"BirthDate", birthdate},
			{"Height", v.Height},
			{"SuperPower", v.SuperPower},
			{"Alive", v.Alive},
		})

		if err != nil {
			log.Fatal(err)
		}

	}
}