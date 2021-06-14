package elastic

import (
	"github.com/elastic/go-elasticsearch/v8"
)

type Repository struct{
	Client 	*elasticsearch.Client
}

func New(uri string, username string, password string)(repo *Repository, err error) {
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
	return repo , err
}
