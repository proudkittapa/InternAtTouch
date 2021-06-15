package user

import (
	"context"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/repository/mongodb"
	"log"
)

type Repository struct {
	*mongodb.Repository
}

func New(ctx context.Context, uri string, dbName string, collName string) (repo *Repository, err error) {
	mongoDB, err := mongodb.New(ctx, uri, dbName, collName)
	log.Println("uri", uri)
	log.Println("dbName", dbName)
	if err != nil {
		return nil, err
	}
	return &Repository{mongoDB}, nil
}

