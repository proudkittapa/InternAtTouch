package elastic

import (
	"bytes"
	"encoding/json"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp2/domain"
	"log"
)

func BuildUpdateRequest(t *domain.UpdateStruct) (buf bytes.Buffer, err error) {
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
		return buf , err
	}
	return buf , err
}

func BuildCheckIndexRequest(indexname string)  (buf bytes.Buffer, err error) {
	query := map[string]interface{}{
		"query" : map[string]interface{}{
			"match": map[string]interface{}{
				"_index": indexname,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return buf , err
	}
	return buf , err
}

func BuildCheckIDRequest(id string)  (buf bytes.Buffer, err error) {
	query := map[string]interface{}{
		"query" : map[string]interface{}{
			"match": map[string]interface{}{
				"_id": id,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return buf , err
	}
	return buf , err
}

func BuildCheckNameRequest(name string)  (buf bytes.Buffer, err error) {
	query := map[string]interface{}{
		"query" : map[string]interface{}{
			"match": map[string]interface{}{
				"name": name,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return buf , err
	}
	return buf , err
}

func BuildCheckActualNameRequest(actualName string)  (buf bytes.Buffer, err error) {
	query := map[string]interface{}{
		"query" : map[string]interface{}{
			"match": map[string]interface{}{
				"actual_name": actualName,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return buf , err
	}
	return buf , err
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