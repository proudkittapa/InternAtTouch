package elastic

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func (repo *Repository) checkExistID(ctx context.Context, id string) (bool, error) {
	req := esapi.ExistsRequest{
		Index:      repo.Index,
		DocumentID: id,
	}

	res, err := req.Do(ctx, repo.Client)
	fmt.Println("res :" ,  res)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()
	return true, err
}

func (repo *Repository) CheckExistName(ctx context.Context, name string) (bool, error) {
	buf, err := BuildCheckNameRequest(name)
	if err != nil{
		return false, err
	}
	result, err := repo.query(ctx,buf)
	if result != nil {} // TODO check if exist or not

	return true, err
}

func (repo *Repository) CheckExistActualName(ctx context.Context, actualName string) (bool, error) {
	buf, err := BuildCheckActualNameRequest(actualName)
	if err != nil{
		return false, err
	}
	result, err := repo.query(ctx,buf)
	if result != nil {} // TODO check if exist or not

	return true, err
}

func (repo *Repository) CheckExistIndex(ctx context.Context, Indexname string) (bool, error) {
	buf, err := BuildCheckIndexRequest(Indexname)
	if err != nil{
		return false, err
	}
	result, err := repo.query(ctx,buf)
	if result != nil {} // TODO check if exist or not

	return true, err
}
