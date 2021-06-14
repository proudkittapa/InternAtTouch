package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
)

func (repo *Repository)query(ctx context.Context,buf bytes.Buffer) (map[string]interface{}, error){
	es := repo.Client
	var r  map[string]interface{}
	res, err := es.Search(
		es.Search.WithContext(ctx),
		es.Search.WithIndex("superhero"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	return r, err
}

func (repo *Repository)search(keyword string,ctx context.Context)(map[string]interface{}, error){
	result, err := repo.query(ctx,buildSearchRequest(keyword))
	return result, err
}

func (repo *Repository)view(id string,ctx context.Context)(map[string]interface{}, error){
	result, err := repo.query(ctx,buildViewRequest(id))
	return result, err
}

func (repo *Repository)viewAll(page int, size int,ctx context.Context)(map[string]interface{}, error){
	result, err := repo.query(ctx,buildViewAllRequest(page, size))
	return result, err
}