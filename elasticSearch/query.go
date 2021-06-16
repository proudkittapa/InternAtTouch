package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	goxid "github.com/touchtechnologies-product/xid"
	"log"
	"strings"
	"sync"
)

var initID string
var temp  map[string]interface{}

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

func createDb(es *elasticsearch.Client) error{
	var wg sync.WaitGroup
	var err error
	for _ , title := range SpList{
		wg.Add(1)

		go func(title InsertStruct) {
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
				//TODO change to receive the name from he config
				Index:      "superhero",
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
	return err
}

func buildRequestUpdate(t UpdateStruct) bytes.Buffer {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"doc": map[string]interface{}{
			"name" : t.Name,
			"actual_name" : t.ActualName,
			"actual_lastname" : t.ActualLastName,
			"gender" : t.Gender,
			"super_power" : t.SuperPower,
			"universe" : t.Universe,
			"movies" : t.Movies,
			"enemies" : t.Enemies,
			"family_member" : t.FamilyMember,
			"about" : t.About,
		},

	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	return buf
}

func insert(ctx context.Context, es *elasticsearch.Client, title InsertStruct) error{
	out, err := json.Marshal(title)
	if err != nil {
		panic (err)
	}

	var b strings.Builder
	b.WriteString(string(out))

	// Set up the request object.
	// TODO get the id from app 1 that generates that id for mongo DB
	initID := goxid.New()
	req := esapi.IndexRequest{
		Index:      "superhero",
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
	return err
}

func update(ctx context.Context, es *elasticsearch.Client, title InsertStruct, id string) error{
	out, err := json.Marshal(title)
	if err != nil {
		panic (err)
	}

	var b strings.Builder
	b.WriteString(string(out))
	// Set up the request object.
	req := esapi.UpdateRequest{
		//TODO update to receive the name from config as well
		Index:      "superhero",
		DocumentID: id,
		Body:       strings.NewReader(b.String()),
		Refresh:    "true",
	}


	// Perform the request with the client.
	res, err := req.Do(ctx, es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	return err
}

func update2(ctx context.Context,es *elasticsearch.Client, title UpdateStruct, id string){
	buf := buildRequestUpdate(title)
	res, err := es.Update(
		//TODO update to receive the name from config as well
		"superhero", id, &buf,
		es.Update.WithContext(ctx),
		es.Update.WithPretty())

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
	if err := json.NewDecoder(res.Body).Decode(&temp); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
}

func Delete(ctx context.Context, es *elasticsearch.Client, id string) error{
	req := esapi.DeleteRequest{
		Index:      "superhero",
		DocumentID: id,
		Refresh:    "true",
	}

	// Perform the request with the client.
	res, err := req.Do(ctx, es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	return err
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
	//createDb(es)

	res, err := es.Info()

	// Deserialize the response into a map.
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print(res)
	}

	figure := InsertStruct{"Black Panther", "T'challa", "-", "Male", 218048400, 183, []string{"Speed", "Strength"}, true, "Marvel", []string{"Black Pabther", "The Avengers0"}, []string{"Erik Killmonger"}, []string{"Shuri", "T'Chaka"}, "The king the Wakanda"}
	//figure := UpdateStruct{"c33gl7aciaega3p5gnsg", "Black Panther", "T'challa", "-", "Female", 218048400, 183, []string{"Speed", "Strength"}, true, "Marvel", []string{"Black Pabther", "The Avengers0"}, []string{"Erik Killmonger"}, []string{"Shuri", "T'Chaka"}, "The king the Wakanda"}
	//insert(ctx, es, figure)
	update2(ctx, es, figure,  "c33gl7aciaega3p5gnsg")
	//Delete(ctx, es, "c33fdpaciaeqk9iid4fg")
}

