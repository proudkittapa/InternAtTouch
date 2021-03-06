package elastic

import (
	"context"
)

func (repo *Repository) CheckExistID(ctx context.Context, id string) (bool, error) {
	buf, err := BuildCheckIDRequest(id)
	if err != nil{
		return false, err
	}
	result, err := repo.query(ctx,buf)
	found := int((result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"]).(float64))
	if found == 0 {
		return false, nil
	}
	return true, err
}


func (repo *Repository) CheckExistName(ctx context.Context, name string) (bool, error) {
	buf, err := BuildCheckNameRequest(name)
	if err != nil{
		return false, err
	}
	result, err := repo.query(ctx,buf)
	found := int((result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"]).(float64))
	if found == 0 {
		return false, nil
	}
	return true, err
}

func (repo *Repository) CheckExistActualName(ctx context.Context, actualName string) (bool, error) {
	buf, err := BuildCheckActualNameRequest(actualName)
	if err != nil{
		return false, err
	}
	result, err := repo.query(ctx,buf)
	found := int((result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"]).(float64))
	if found == 0 {
		return false, nil
	}
	return true, err
}

func (repo *Repository) CheckExistIndex(ctx context.Context, Indexname string) (bool, error) {
	buf, err := BuildCheckIndexRequest(Indexname)
	if err != nil{
		return false, err
	}
	result, err := repo.query(ctx,buf)
	found := int((result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"]).(float64))
	if found == 0 {
		return false, nil
	}
	return true, err
}
