package elastic

import (
	"github.com/elastic/go-elasticsearch/v8"
)

type Repository struct{
	Client 	*elasticsearch.Client
	Index 	string
}

func New(uri string, username string, password string, indexName string)(repo *Repository, err error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			uri,
	},
		Username : username,
		Password : password,
	}

	es, err := elasticsearch.NewClient(cfg)
	repo = &Repository{}
	repo.Client = es
	repo.Index = indexName
	return repo , err
}
