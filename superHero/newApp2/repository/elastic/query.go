package elastic

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/domain"
	"log"
	"strings"
)

func (repo *Repository) Insert(ctx context.Context, title *domain.UpdateStruct) error{
	out, err := json.Marshal(title)
	if err != nil {
		return err
	}

	var b strings.Builder
	b.WriteString(string(out))

	// TODO get the id from app 1 that generates that id for mongo DB
	req := esapi.IndexRequest{
		Index:      "superhero",
		DocumentID: title.ID,
		Body:       strings.NewReader(b.String()),
		Refresh:    "true",
	}

	res, err := req.Do(ctx, repo.Client)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return err
}

func (repo *Repository)Update(ctx context.Context, title *domain.UpdateStruct) error{
	buf := BuildUpdateRequest(title)
	res, err := repo.Client.Update(
		//TODO update to receive the name from config as well
		"superhero", title.ID, &buf,
		repo.Client.Update.WithContext(ctx),
		repo.Client.Update.WithPretty())

	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return err
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	return err
}

func (repo *Repository)Delete(ctx context.Context, id string) error{
	req := esapi.DeleteRequest{
		Index:      "superhero",
		DocumentID: id,
		Refresh:    "true",
	}

	res, err := req.Do(ctx, repo.Client)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	return err
}