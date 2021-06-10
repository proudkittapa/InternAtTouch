package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"touch/Database"
	"touch/validatePack"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Pagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type Person struct {
	Name string
}

var validate *validator.Validate

func main() {
	validate = validator.New()

	// validate.RegisterValidation("updateName", updateName)
	// validate.RegisterValidation("updateActualName", updateActualName)
	validate.RegisterStructValidation(existanceActualName, Database.SuperheroQ{})
	validate.RegisterStructValidation(existanceName, Database.SuperheroQ{})
	validate.RegisterStructValidation(updateName, Database.UpdateSuperhero{})
	validate.RegisterStructValidation(updateActualName, Database.UpdateSuperhero{})
	validate.RegisterStructValidation(updateID, Database.UpdateSuperhero{})

	r := setupRouter()
	Database.InitDB()
	r.Run()

}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/superheroes", insert)
	r.PUT("/superheroes/:id", updateId)
	r.DELETE("/superheroes/:id", deleteId)
	r.GET("/superheroes/:id", viewId)
	r.GET("/superheroes", viewAll)
	r.GET("/superheroes/search", search)

	return r
}

func insert(c *gin.Context) {
	var hero Database.SuperheroQ
	if err := c.ShouldBindJSON(&hero); err != nil {
		c.JSON(http.StatusBadRequest, "can't bind, check the format")
		return
	}
	// if Database.CheckExistName(user.Name) { //if jer
	// 	c.JSON(http.StatusBadRequest, "name already exist")
	// 	return
	// }
	val, message := validateInsert(hero)
	fmt.Println("hrerhere", message)
	if !val {
		c.JSON(http.StatusBadRequest, message)
		return
	}
	// Database.Insert(user)
	c.JSON(http.StatusOK, "inserted")
}

func updateId(c *gin.Context) {
	id := c.Param("id")

	// if !Database.CheckExistID(id) {
	// 	c.JSON(http.StatusNotFound, "this id doesn't exist")
	// 	return
	// }
	var hero Database.UpdateSuperhero
	if err := c.ShouldBindJSON(&hero); err != nil {
		c.JSON(http.StatusBadRequest, "can't bind")
		return
	}

	hero.ID = id

	val, message := validateUpdate(hero)
	log.Println(message)
	if !val {
		c.JSON(http.StatusBadRequest, message)
		return
	}

	// Database.Update(user, id)
	c.JSON(http.StatusOK, "updated")
}

func deleteId(c *gin.Context) {
	id := c.Param("id")
	// if !validatePack.CheckExistID(id) {
	// 	// fmt.Println("this id doesn't exist")
	// 	c.JSON(http.StatusNotFound, "this id doesn't exist")
	// 	return
	// }
	Database.Delete(id)
	c.JSON(http.StatusOK, "deleted")
}

func viewId(c *gin.Context) {
	id := c.Param("id")

	// if !validatePack.CheckExistID(id) {
	// 	// fmt.Println("this id doesn't exist")
	// 	c.JSON(http.StatusNotFound, "this id doesn't exist")
	// 	return
	// }

	var data Database.SuperheroQ = Database.View(id)
	c.JSON(http.StatusOK, data)
}

func viewAll(c *gin.Context) {
	p := pagination(c)
	fmt.Println(p)
	data := Database.ViewByPage(p.Limit, p.Page)
	if data == nil {
		c.JSON(http.StatusNotFound, "this page is not available")
		return
	}
	c.JSON(http.StatusOK, data)
}

func pagination(c *gin.Context) Pagination {
	limit := 2
	page := 0
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		}
	}
	return Pagination{
		Limit: limit,
		Page:  page,
	}
}

func search(c *gin.Context) {
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	fmt.Println(reqBody)
	var v Database.SearchValue
	err := json.Unmarshal(buf[0:num], &v)
	if err != nil {
		fmt.Println(err)
	}
	data := Database.Search(v.Value)
	if data == nil {
		c.JSON(http.StatusOK, "No result")
		return
	}
	c.JSON(http.StatusOK, data)
}

type Err struct {
	Code  int
	Cause []string
}

type Name struct {
	n string
}

func validateInsert(hero Database.SuperheroQ) (b bool, out Err) {
	err := validate.Struct(hero)
	var arr []string
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			// name := Name{
			// 	n: string(err.StructField()) + " " + string(err.ActualTag()),
			// }
			name := string(err.StructField()) + " " + string(err.ActualTag())
			arr = append(arr, name)
			out = Err{
				Code:  400,
				Cause: arr,
			}
		}
		log.Println("out:", out)

		b = false
		return b, out
	}
	b = true
	return b, out
}
func validateUpdate(hero Database.UpdateSuperhero) (b bool, out Err) {
	fmt.Println("user:", hero)
	err := validate.Struct(hero)
	fmt.Println(err)
	var arr []string
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			// name := Name{
			// 	n: string(err.StructField()) + " " + string(err.ActualTag()),
			// }
			name := string(err.StructField()) + " " + string(err.ActualTag())
			arr = append(arr, name)
			out = Err{
				Code:  400,
				Cause: arr,
			}
		}
		log.Println("update out", out)

		b = false
		return b, out
	}
	b = true
	return b, out
}

func existanceActualName(structLV validator.StructLevel) {
	input := structLV.Current().Interface().(Database.SuperheroQ)
	validatePack.CheckExistActualName(structLV, input)
}
func existanceName(structLV validator.StructLevel) {
	input := structLV.Current().Interface().(Database.SuperheroQ)
	validatePack.CheckExistName(structLV, input)
}

func updateActualName(structLV validator.StructLevel) {
	fmt.Println("herejrhe")
	input := structLV.Current().Interface().(Database.UpdateSuperhero)
	validatePack.CheckUpdateActualName(structLV, input)
}

func updateName(structLV validator.StructLevel) {
	input := structLV.Current().Interface().(Database.UpdateSuperhero)
	validatePack.CheckUpdateName(structLV, input)
}
func updateID(structLV validator.StructLevel) {
	input := structLV.Current().Interface().(Database.UpdateSuperhero)
	validatePack.CheckExistID(structLV, input)
}
