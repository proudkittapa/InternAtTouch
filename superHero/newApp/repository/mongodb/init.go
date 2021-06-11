package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct{
	Client 	*mongo.Client
	DB    	*mongo.Database
	Coll	*mongo.Collection
	URI    	string
	DBName 	string
}

func New(ctx context.Context , uri string, dbName string, collName string)(repo *Repository, err error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	repo = &Repository{}
	repo.URI = uri
	repo.DBName = dbName
	repo.Client = client
	repo.DB = client.Database(dbName)
	repo.Coll = repo.DB.Collection(collName)

	return repo, nil
}
