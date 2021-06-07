package Search

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"log"
	"time"
	"touch/Database"
)

var collection *mongo.Collection
var ctx context.Context

func init() {
	uri := "mongodb://touch:touchja@localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("superheroes").Collection("list")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

}

func Search(keyword string) []Database.Superhero_q {
	fmt.Println("Searching")
	var result []Database.Superhero_q
	cursor, err := collection.Find(ctx, bson.M{"Name": primitive.Regex{Pattern: "^" + keyword + ".*", Options: "i"}})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {

		var result_bson bson.M
		var result_struct Database.Superhero_q
		if err = cursor.Decode(&result_bson); err != nil {
			log.Fatal(err)
		}
		bsonBytes, _ := bson.Marshal(result_bson)
		bson.Unmarshal(bsonBytes, &result_struct)
		fmt.Println(result_struct)
		result = append(result, result_struct)
	}
	return result
}
