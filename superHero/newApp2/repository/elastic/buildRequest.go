package elastic

import (
	"bytes"
	"encoding/json"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp2/domain"
	"log"
)

func BuildUpdateRequest(t domain.UpdateStruct) bytes.Buffer {
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