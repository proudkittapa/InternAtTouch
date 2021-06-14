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

type ElasticDocs struct {
	Name      		string   `bson:"name" json:"name" validate:"required"`
	ActualName 		string   `bson:"actual_name" json:"actual_name"`
	ActualLastName  string   `bson:"actual_lastname" json:"actual_lastname"`
	Gender     		string   `bson:"gender" json:"gender"`
	BirthDate  		int64    `bson:"birth_date" json:"birth_date"`
	Height     		int      `bson:"height" json:"height" validate:"gte=0"`
	SuperPower 		[]string `bson:"super_power" json:"super_power"`
	Alive      		bool     `bson:"alive" json:"alive"`
	Universe 		string	 `bson:"universe" json:"universe"`
	Movies			[]string `bson:"movies" json:"movies"`
	Enemies			[]string `bson:"enemies" json:"enemies"`
	FamilyMember	[]string `bson:"family_member" json:"family_member"`
	About			string	 `bson:"about" json:"about"`
}

var spList =  []ElasticDocs{
	{"Spider-Man", "Peter", "Parker", "Male", 997401600, 178, []string{"Web-shooting"}, true, "Marvel", []string{"Spiderman", "The Avengers"}, []string{"Globlin", "Doctor Octopus"}, []string{"Richard Parker", "Mary Parker"}, "A boy who has been bitten by a spider and become superhero."},
	{"Batman", "Bruce", "Wayne", "Male", 261619200, 188, []string{"Rich"}, true, "DC", []string{"Batman", "Justice League", "The Dark Knight"}, []string{"Joker", "Superman"}, []string{"Tim Drake", "Cassandra Cain"}, "A rich man who want to be a superhero."},
	{"Superman", "Clark", "Kent", "Male", 230169600, 191, []string{"Flight", "Strength"}, true, "DC", []string{"Superman", "Man of Steel", "Justice League"}, []string{"Batman", "Justice league"}, []string{"Kara Kent", "Linda Danvers"}, "A alien who come from Krypton and become superhero in the earth."},
	{"Wonder woman", "Diana", "Prince", "Female", -908236800, 178, []string{"Agility" , "Strength"}, true, "DC", []string{"Wonder Woman", "Justice League"}, []string{"Doctor Poison"}, []string{"Donna Troy", "Miss America"}, "A girl from the mystery island and become a superhero in the real world."},
	{"Doctor Strange", "Stephen", "Strange", "Male", -1234569600, 183, []string{"Magic"}, true, "Marvel", []string{"Doctor Strange", "The Avengers"}, []string{"Baron Karl Amadeus Mordo", "Thanos"}, []string{"Donna Strange"}, "A doctor who has a lot of maditation and become a superhero."},
	{"Iron man", "Tony", "Stark", "Male", 12787200, 185, []string{"Genius", "super-suit"}, false, "Marvel", []string{"Iron Man", "The Avengers"}, []string{"Mandarin", "Doctor Doom"}, []string{"Howard Stark", "Maria Stark"}, "A rich man who become superhero in the super suit."},
	{"Black Widow", "Natasha", "Romanoff", "Female", 469929600, 170, []string{"Expert spy"}, false, "Marvel", []string{"Captain America", "The Avengers"}, []string{"Thanos", "Hulk"}, []string{"Melina Vostokoff", "Yelena Belova"}, "A Shield's spy agent who become a superhero."},
	{"Scarlet Witch", "Wanda", "Maximoff",  "Female", 192758400, 170, []string{"Energy manipulation"}, true, "Marvel", []string{"Captain America", "The Avengers"}, []string{"Iron man", "Ultron"}, []string{"Marya Maximoff", "Natalya Maximoff"}, "A witch who become superhero."},
	{"Harley Quinn", "Harleen", "Quinzel", "Female",  929836800, 168, []string{"Immunity", "Strength"}, true, "DC", []string{"Suicide Squad", "Birds of Prey"}, []string{"Batman", "Brimstone"}, []string{"Delia Quinn"}, "A witch who become superhero."},
	{"Captain America", "Steve", "Rogers", "Male", -1625097600, 188, []string{"Immunity", "Strength"}, true, "Marvel", []string{"Captain America", "The Avengers"}, []string{"Red Skull", "Thanos"}, []string{"Michael Rogers"}, "A witch who become superhero."},
}
// A function for marshaling structs to JSON string
func jsonStruct(doc ElasticDocs) string {

	// Create struct instance of the Elasticsearch fields struct object
	docStruct := &ElasticDocs{
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

	for _, hero := range spList{
		heroStr := jsonStruct(hero)
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
