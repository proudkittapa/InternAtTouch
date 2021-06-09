package Database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbconfig struct{
	Coll	*mongo.Collection
	Ctx		context.Context
}


func InitDB() {
	var heroDB dbconfig
	uri := os.Getenv("MONGODB_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	heroDB.Coll = client.Database("superhero").Collection("lists")
	heroDB.Ctx,_ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(heroDB.Ctx)
}
