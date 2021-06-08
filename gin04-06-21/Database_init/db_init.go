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
	Name       string   `bson:"name" json:"name" validate:"required"`
	ActualName string   `bson:"actual_name" json:"actual_name"`
	Gender     string   `bson:"gender" json:"gender"`
	BirthDate  int64    `bson:"birth_date" json:"birth_date"`
	Height     int      `bson:"height" json:"height" validate:"gte=0"`
	SuperPower []string `bson:"super_power" json:"super_power"`
	Alive      bool     `bson:"alive" json:"alive"`
}

var Sp_list =  []Sp{
	{"Spider-Man", "Peter Parker", "Male", 997401600, 178, []string{"Web-shooting"}, true},
	{"Batman", "Bruce Wayne", "Male", 261619200, 188, []string{"Rich"}, true},
	{"Superman", "Clark Kent", "Male", 230169600, 191, []string{"Flight", "Strength"}, true},
	{"Wonder woman", "Diana Prince", "Female", -908236800, 178, []string{"Agility" , "Strength"}, true},
	{"Doctor Strange", "Stephen Vincent Strange", "Male", -1234569600, 183, []string{"Magic"}, true},
	{"Iron man", "Tony Stark", "Male", 12787200, 185, []string{"Genius", "super-suit"}, false},
	{"Black Widow", "Natasha Romanoff", "Female", 469929600, 170, []string{"Expert spy"}, false},
	{"Scarlet Witch", "Wanda Maximoff", "Female", 192758400, 170, []string{"Energy manipulation"}, true},
	{"Harley Quinn", "Dr. Harleen Quinzel", "Female",  929836800, 168, []string{"Immunity and Strength"}, true},
	{"Captain America", "Steve Rogers", "Male", -1625097600, 188, []string{"Immunity", "Strength"}, true},
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
			{"BirthDate", v.BirthDate},
			{"Height", v.Height},
			{"SuperPower", v.SuperPower},
			{"Alive", v.Alive},
		})

		if err != nil {
			log.Fatal(err)
		}

	}
	_, err = collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.M{
				"Name": 1,
			},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(indexName)
	_, err = collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.M{
				"ActualName": 1,
			},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(indexActualName)
}