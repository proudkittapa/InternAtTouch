package elastic

import (
	"bytes"
	//"context"
	"encoding/json"
	//"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	//"reflect"
	"strconv"
	//"strings"
	//"sync"
)

func buildSearchRequest(keyword string) bytes.Buffer {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"query_string": map[string]interface{}{
				"query" : "*"+keyword+"*",
				"fields" : []interface{}{
					"name",
					"actual_name",
					"actual_lastname",
					"gender",
					"super_power",
					"universe",
					"movies",
					"enemies",
					"family_member",
					"about",
				},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
	log.Fatalf("Error encoding query: %s", err)
	}
	return buf
}

func buildViewAllRequest(page int, size int) bytes.Buffer{
	var buf bytes.Buffer
	page = strconv.Atoi(page)
	size = strconv.Atoi(size)
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

func buildViewRequest(id string) bytes.Buffer{
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query" : map[string]interface{}{
			"match": map[string]interface{}{
				"id": id,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	return buf
}