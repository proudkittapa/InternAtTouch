package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"reflect"
	"strconv"
	"strings"
)

// A function for marshaling structs to JSON string
func jsonStructInit(doc Sp) string {

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

	//fmt.Println("\ndocStruct:", docStruct)
	//fmt.Println("docStruct TYPE:", reflect.TypeOf(docStruct))

	// Marshal the struct to JSON and check for errors
	b, err := json.Marshal(docStruct)
	if err != nil {
		fmt.Println("json.Marshal ERROR:", err)
		return string(err.Error())
	}
	return string(b)
}

func main(){
	log.SetFlags(0)

	// Create a context object for the API calls
	ctx := context.Background()

	// Create a mapping for the Elasticsearch documents
	var (
		docMap map[string]interface{}
	)
	fmt.Println("docMap:", docMap)
	fmt.Println("docMap TYPE:", reflect.TypeOf(docMap))

	// Declare an Elasticsearch configuration
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Username: "user",
		Password: "pass",
	}

	// Instantiate a new Elasticsearch client object instance
	client, err := elasticsearch.NewClient(cfg)
	createDb(client)
	if err != nil {
		fmt.Println("Elasticsearch connection error:", err)
	}

	// Have the client instance return a response
	res, err := client.Info()

	// Deserialize the response into a map.
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print(res)
	}

	// Declare empty array for the document strings
	var docs []string
	//
	//// Declare documents to be indexed using struct
	//doc1 := ElasticDocs{}
	//doc1.SomeStr = "Some Value"
	//doc1.SomeInt = 123456
	//doc1.SomeBool = true
	//
	//doc2 := ElasticDocs{}
	//doc2.SomeStr = "Another Value"
	//doc2.SomeInt = 42
	//doc2.SomeBool = false
	//
	//// Marshal Elasticsearch document struct objects to JSON string
	//docStr1 := jsonStruct(doc1)
	//docStr2 := jsonStruct(doc2)
	//
	//// Append the doc strings to an array
	//docs = append(docs, docStr1)
	//docs = append(docs, docStr2)

	for _, hero := range SpList{
		heroStr := jsonStructInit(hero)
		docs = append(docs, heroStr)
	}

	// Iterate the array of string documents
	for i, bod := range docs {
		//fmt.Println("\nDOC _id:", i+1)
		//fmt.Println(bod)

		// Instantiate a request object
		req := esapi.IndexRequest{
			Index:      "name",
			DocumentID: strconv.Itoa(i + 1),
			Body:       strings.NewReader(bod),
			Refresh:    "true",
		}
		fmt.Println(reflect.TypeOf(req))

		// Return an API response object from request
		res, err := req.Do(ctx, client)
		if err != nil {
			log.Fatalf("IndexRequest ERROR: %s", err)
		}
		//defer res.Body.Close()

		if res.IsError() {
			log.Printf("%s ERROR indexing document ID=%d", res.Status(), i+1)
		} else {

			// Deserialize the response into a map.
			var resMap map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
				log.Printf("Error parsing the response body: %s", err)
			} else {
				//log.Printf("\nIndexRequest() RESPONSE:")
				// Print the response status and indexed document version.
				//fmt.Println("Status:", res.Status())
				//fmt.Println("Result:", resMap["result"])
				//fmt.Println("Version:", int(resMap["_version"].(float64)))
				fmt.Println("resMap:", resMap)
				//fmt.Println("\n")
			}
		}
	}
}
