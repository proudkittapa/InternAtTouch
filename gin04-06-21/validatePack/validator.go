package validatePack

import (
	"fmt"
	"log"
	"touch/Database"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckExistID(structLV validator.StructLevel, input Database.UpdateSuperhero) {
	count, err := Database.Coll.CountDocuments(Database.Ctx, bson.D{{"_id", input.ID}})
	if err != nil {
		log.Fatal("err1: ", err)
	}
	if count < 1 {
		structLV.ReportError("ID is not in the database", "id", "id", "unique", "")
	}
}

func checkExistName(name string) bool {
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

func checkExistActualName(actualName string) bool {
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

func CheckUpdateActualName(structLV validator.StructLevel, input Database.UpdateSuperhero) {
	if !checkExistActualName(input.ActualName) {
		checkActName := Database.View(input.ID)
		fmt.Println(checkActName, input)
		if checkActName.ActualName != input.ActualName {
			// log.Println("update:")
			structLV.ReportError("same actual name, but not same id", "actual_name", "actual_name", "unique", "")
		}
	}
	// return true
}
func CheckUpdateName(structLV validator.StructLevel, input Database.UpdateSuperhero) {
	if !checkExistName(input.Name) {
		checkName := Database.View(input.ID)
		fmt.Println(checkName, input)
		if checkName.Name != input.Name {
			// log.Println("update:")
			structLV.ReportError("same name, but not same id", "actual_name", "actual_name", "unique", "")
		}
	}
	// return true
}

func CheckExistName(structLV validator.StructLevel, input Database.SuperheroQ) {
	count, err := Database.Coll.CountDocuments(Database.Ctx, bson.D{{"name", input.Name}})
	if err != nil {
		log.Fatal("err2: ", err)
	}
	if count >= 1 {
		structLV.ReportError("name already existed", "name", "name", "unique", "")
	}
}

func CheckExistActualName(structLV validator.StructLevel, input Database.SuperheroQ) {
	count, err := Database.Coll.CountDocuments(Database.Ctx, bson.D{{"actual_name", input.ActualName}})
	if err != nil {
		log.Fatal("err2: ", err)
	}
	if count >= 1 {
		structLV.ReportError("actual name already existed", "actual_name", "actual_name", "unique", "")
	}
}
