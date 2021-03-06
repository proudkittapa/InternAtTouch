package Database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Dbconfig struct{
	Client 	*mongo.Client
	DB    	*mongo.Database
	Coll	*mongo.Collection
	Ctx		context.Context
	URI    	string
	DBName 	string
}

func InitDB(uri string, dbName string, collName string)(repo *Dbconfig, err error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	repo = &Dbconfig{}
	repo.URI = uri
	repo.DBName = dbName
	repo.Client = client
	repo.DB = client.Database(dbName)
	repo.Coll = repo.DB.Collection(collName)
	repo.Ctx,_ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(repo.Ctx)
	if err != nil {
		return nil, err
	}
	return repo, nil
}
