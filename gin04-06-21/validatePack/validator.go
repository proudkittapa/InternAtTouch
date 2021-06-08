package validatePack

import (
	"fmt"
	"log"
	"touch/Database"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckExistID(id string) bool {
	count, err := Database.Coll.CountDocuments(Database.Ctx, bson.D{{"_id", id}})
	if err != nil {
		log.Fatal("err1: ", err)
		return true
	}
	if count >= 1 {
		return true
	}
	return false
}

func CheckExistName(name string) bool {
	count, err := Database.Coll.CountDocuments(Database.Ctx, bson.D{{"name", name}})
	if err != nil {
		log.Fatal("err2: ", err)
		return false
	}
	if count >= 1 {
		return false
	}
	return true
}

func CheckExistActualName(actualName string) bool {
	count, err := Database.Coll.CountDocuments(Database.Ctx, bson.D{{"actual_name", actualName}})
	if err != nil {
		log.Fatal("err2: ", err)
		return false
	}
	if count >= 1 {
		return false
	}
	return true
}

func CheckUpdateActualName2(structLV validator.StructLevel, input Database.UpdateSuperhero) {
	if !CheckExistActualName(input.ActualName) {
		checkActName := Database.View(input.ID)
		fmt.Println(checkActName, input)
		if checkActName.ActualName != input.ActualName {
			fmt.Println("erororororororo")
			structLV.ReportError("same actual name, but not same id", "license", "license", "unique", "")
		}
	}
	// return true
}

func CheckUpdateName(name string, id string) bool {
	if !CheckExistName(name) {
		checkName := Database.View(id)
		if checkName.Name == name {
			return true
		}
		return false
	} else {
		return true
	}
}

// func CheckUpdateActualName(actualName string, id string) bool {
// 	if !CheckExistActualName(actualName) {
// 		checkActName := View(id)
// 		if checkActName.ActualName == actualName {
// 			return true
// 		}
// 		return false
// 	} else {
// 		return true
// 	}
// }
