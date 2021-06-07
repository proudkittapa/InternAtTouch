package Database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Coll *mongo.Collection
var Ctx context.Context

func InitDB() {
	uri := os.Getenv("MONGODB_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	Coll = client.Database("superheroes").Collection("list")
	Ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(Ctx)
}
