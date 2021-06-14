package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	goxid "github.com/touchtechnologies-product/xid"
	"log"
	"strings"
	"sync"
)

func initDb(uri string, username string, password string)(*elasticsearch.Client, error){
	cfg := elasticsearch.Config{
		Addresses: []string{
			uri,
		},
		Username: username,
		Password: password,
	}
	es, err := elasticsearch.NewClient(cfg)

	return es,err
}

func createDb(es *elasticsearch.Client){
	var wg sync.WaitGroup
	for _ , title := range SpList{
		wg.Add(1)

		go func(title Sp) {
			defer wg.Done()
			out, err := json.Marshal(title)
			if err != nil {
				panic (err)
			}

			var b strings.Builder
			b.WriteString(string(out))

			// Set up the request object.
			initID := goxid.New()
			req := esapi.IndexRequest{
				Index:      "list",
				DocumentID: initID.Gen(),
				Body:       strings.NewReader(b.String()),
				Refresh:    "true",
			}

			// Perform the request with the client.
			res, err := req.Do(context.Background(), es)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}
			defer res.Body.Close()
		}(title)
	}
	wg.Wait()
}

func structToJson(doc Sp) string {
	// Create struct instance of the Elasticsearch fields struct object
	docStruct := &Sp{
		Name: doc.Name,
		ActualName: doc.ActualName,
		ActualLastName: doc.ActualLastName,
		Gender: doc.Gender,
		BirthDate: doc.BirthDate,
		Height: doc.Height,
		SuperPower: doc.SuperPower,
		Alive: doc.Alive,
		Universe: doc.Universe,
		Movies: doc.Movies,
		Enemies: doc.Enemies,
		FamilyMember: doc.FamilyMember,
		About: doc.About,
	}

	// Marshal the struct to JSON and check for errors
	b, err := json.Marshal(docStruct)
	if err != nil {
		fmt.Println("json.Marshal ERROR:", err)
		return string(err.Error())
	}
	return string(b)
}

func upsert(ctx context.Context, es *elasticsearch.Client, title Sp){
	out, err := json.Marshal(title)
	if err != nil {
		panic (err)
	}

	var b strings.Builder
	b.WriteString(string(out))

	// Set up the request object.
	initID := goxid.New()
	req := esapi.IndexRequest{
		Index:      "list",
		DocumentID: initID.Gen(),
		Body:       strings.NewReader(b.String()),
		Refresh:    "true",
	}


	// Perform the request with the client.
	res, err := req.Do(ctx, es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
}

func buildQueryID(keyword string) bytes.Buffer {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"_id" : keyword,
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	return buf
}

func Create(ctx context.Context,es *elasticsearch.Client, res *esapi.Response, buf bytes.Buffer, err error){
	res, err = es.Search(
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
			// Print the response status and error information.
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
}

func main(){
	log.SetFlags(0)
	ctx := context.Background()

	var (
		//r map[string]interface{}
	)

	es , err := initDb("http://localhost:9200", "touch", "touchja")
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	createDb(es)

	figure := Sp{"Black Panther", "T'challa", "-", "Male", 218048400, 183, []string{"Speed", "Strength"}, true, "Marvel", []string{"Black Pabther", "The Avengers0"}, []string{"Erik Killmonger"}, []string{"Shuri", "T'Chaka"}, "The king the Wakanda"}
	upsert(ctx, es, figure)



}

