package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"reflect"
	//"strconv"

	//"strconv"
	"strings"
	//"sync"
)

//var r  map[string]interface{}


//func buildRequest(keyword string) bytes.Buffer {
//	var buf bytes.Buffer
//	query := map[string]interface{}{
//		"query": map[string]interface{}{
//			"query_string": map[string]interface{}{
//				"query" : "*"+keyword+"*",
//				"fields" : []interface{}{
//					"name",
//					"actual_name",
//					"actual_lastname",
//					"gender",
//					"super_power",
//					"universe",
//					"movies",
//					"enemies",
//					"family_member",
//					"about",
//				},
//				},
//			},
//	}
//	if err := json.NewEncoder(&buf).Encode(query); err != nil {
//		log.Fatalf("Error encoding query: %s", err)
//	}
//	return buf
//}

//func buildRequest(keyword string) bytes.Buffer {
//	var buf bytes.Buffer
//	query := map[string]interface{}{
//		"from": "0",
//		"size": "20",
//		"query" : map[string]interface{}{
//			"match": map[string]interface{}{
//				"_index": "superhero",
//				},
//			},
//		}
//
//
//	if err := json.NewEncoder(&buf).Encode(query); err != nil {
//		log.Fatalf("Error encoding query: %s", err)
//	}
//	return buf
//}

func buildRequest(page int, size int) bytes.Buffer{
	var buf bytes.Buffer
	from := (page-1)*size
	query := map[string]interface{}{
		"from": from,
		"size": size,
		"query" : map[string]interface{}{
			"match": map[string]interface{}{
				"_index": "superhero",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	return buf
}
//
func search(ctx context.Context,es *elasticsearch.Client, res *esapi.Response, buf bytes.Buffer, err error){
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
	s := res.String()
a := strings.Split(s,"hits")
fmt.Println(a[2])

	//json.Unmarshal([]byte(s), &data)
	//fmt.Println(data)
	//fmt.Println(res.String())
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

}

//func add(ctx context.Context,es *elasticsearch.Client, res *esapi.Response, buf bytes.Buffer, err error){
//	res, err = es.Create(
//		es.Create.WithContext(ctx),
//		es.Create.WithIndex("superhero"),
//		es.Create.WithBody(&buf),
//		es.Create.WithTrackTotalHits(true),
//		es.Create.WithPretty(),
//	)
//	if err != nil {
//		log.Fatalf("Error getting response: %s", err)
//	}
//	defer res.Body.Close()
//}

func main() {

	// Allow for custom formatting of log output
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
	//createDb(client)
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

	// Declare empty array for the document string
	log.Println(strings.Repeat("=", 37))
	search(ctx, client, res, buildRequest(1, 10), err)
	//log.Printf(
	//	"[%s] %d hits; took: %dms",
	//	res.Status(),
	//	int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
	//	int(r["took"].(float64)),
	//)
	// Print the ID and document source for each hit.
	//
	//var temp Sp
	//var ans []Sp
	//for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
	//
	//	s := hit.(map[string]interface{})["_source"]
	//	log.Println("type :",reflect.TypeOf(s.(map[string]interface{})["birth_date"]))
	//	xxx := s.(map[string]interface{})["birth_date"])
	//	temp.Name = fmt.Sprintf("%v", s.(map[string]interface{})["name"])
	//	temp.ActualName = fmt.Sprintf("%v", s.(map[string]interface{})["actual_name"])
	//	temp.ActualLastName = fmt.Sprintf("%v", s.(map[string]interface{})["actual_lastname"])
	//	temp.Gender = fmt.Sprintf("%v", s.(map[string]interface{})["gender"])
	//	temp.BirthDate = int64(xxx)
	//	temp.Height,_ = strconv.Atoi(fmt.Sprintf("%v", s.(map[string]interface{})["height"]))
	//	temp.SuperPower = strings.Split(fmt.Sprintf("%v", s.(map[string]interface{})["super_power"]),",")
	//	temp.Alive,_ = strconv.ParseBool(fmt.Sprintf("%v", s.(map[string]interface{})["alive"]))
	//	temp.Universe = fmt.Sprintf("%v", s.(map[string]interface{})["universe"])
	//	temp.Movies = strings.Split(fmt.Sprintf("%v", s.(map[string]interface{})["movies"]),",")
	//	temp.Enemies = strings.Split(fmt.Sprintf("%v", s.(map[string]interface{})["enemies"]),",")
	//	temp.FamilyMember = strings.Split(fmt.Sprintf("%v", s.(map[string]interface{})["family_member"]),",")
	//	temp.About = fmt.Sprintf("%v", s.(map[string]interface{})["about"])
	//	//x := (s.(map[string]interface{})["birth_date"]).(int64)
	//	//fmt.Println(x)
	//	ans = append(ans, temp)
	//	}
	//	for _, s := range ans{
	//		fmt.Println(s)
	//	}
	//}
	//for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
	//
	//	for _, s := range hit.(map[string]interface{})["_source"].(map[string]interface{}) {
	//		str := fmt.Sprintf("%v", s)
	//		fmt.Println("pond",str)
	//	}
	//
	//
	//
	//}
	fmt.Println(r)

}