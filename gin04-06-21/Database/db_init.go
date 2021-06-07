package Database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Sp struct {
	Name  		string	`bson:"Name"`
	Actual_name string	`bson:"Actual_name"`
	Gender 		string	`bson:"Gender"`
	Age 		int		`bson:"Age"`
	Super_power string	`bson:"Super_power"`
}
var Sp_list =  []Sp{
	{"Spider-Man", "Peter Parker", "Male", 28, "Web-shooting"},
	{"Batman", "Bruce Wayne", "Male", 75, "Rich"},
	{"Superman", "Clark Kent", "Male", 79, "Flight and Strength"},
	{"Wonder woman", "Diana Prince", "Female", 800, "Agility and Strength"},
	{"Doctor Strange", "Stephen Vincent Strange", "Male", 36, "Magic"},
	{"Iron man", "Tony Stark", "Male", 48, "Genius and super-suit"},
	{"Black Widow", "Natasha Romanoff", "Female", 33, "Expert spy"},
	{"Scarlet Witch", "Wanda Maximoff", "Female", 20, "Energy manipulation"},
	{"Harley Quinn", "Dr. Harleen Quinzel", "Female", 25, "Immunity and Strength"},
	{"Captain America", "Steve Rogers", "Male", 93, "Immunity and Strength"},
}

func Db_init(){

	uri := "mongodb://touch:touchja@localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	collection := client.Database("superheroes").Collection("list")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}


	defer client.Disconnect(ctx)
	for k ,v := range Sp_list{
		//g := &Sp{v.Name, v.Actual_name, v.Gender, v.Age, v.Super_power}
		//fmt.Println(v.Name, v.Actual_name, v.Gender, v.Age, v.Super_power)
		//_, err = collection.InsertOne(ctx, g)
		_, err = collection.InsertOne(ctx, bson.D{
			{"ID", k+1},
			{"Name", v.Name},
			{"Actual_name", v.Actual_name},
			{"Gender", v.Gender},
			{"Age", v.Age},
			{"Super_power", v.Super_power},
		})

		if err != nil {
			log.Fatal(err)
		}

	}
}