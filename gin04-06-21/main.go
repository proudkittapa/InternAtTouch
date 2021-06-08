package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"touch/Database"

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
	validate.RegisterValidation("uniqueActualName", existanceActualName)
	validate.RegisterValidation("uniqueName", existanceName)
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
		c.JSON(http.StatusBadRequest, "can't bind")
		return
	}
	if Database.CheckExistName(hero.Name) { //if jer
		c.JSON(http.StatusBadRequest, "name already exist")
		return
	}
	val, message := validateHero(hero)
	if !val {
		c.JSON(http.StatusBadRequest, message)
		return
	}
	Database.Insert(hero)
	c.JSON(http.StatusOK, "inserted")
}

func updateId(c *gin.Context) {
	id := c.Param("id")

	// if !Database.CheckExistID(id) {
	// 	c.JSON(http.StatusNotFound, "this id doesn't exist")
	// 	return
	// }
	var hero Database.SuperheroQ
	if err := c.ShouldBindJSON(&hero); err != nil {
		c.JSON(http.StatusBadRequest, "can't bind")
		return
	}

	val, message := validateHero(hero)
	if !val {
		c.JSON(http.StatusBadRequest, message)
		return
	}

	Database.Update(hero, id)
	c.JSON(http.StatusOK, "updated")
}

func deleteId(c *gin.Context) {
	id := c.Param("id")
	if !Database.CheckExistID(id) {
		// fmt.Println("this id doesn't exist")
		c.JSON(http.StatusNotFound, "this id doesn't exist")
		return
	}
	Database.Delete(id)
	c.JSON(http.StatusOK, "deleted")
}

func viewId(c *gin.Context) {
	id := c.Param("id")

	if !Database.CheckExistID(id) {
		// fmt.Println("this id doesn't exist")
		c.JSON(http.StatusNotFound, "this id doesn't exist")
		return
	}

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
	data := Database.SearchContainName(v.Value)
	if data == nil {
		c.JSON(http.StatusOK, "No result")
		return
	}
	c.JSON(http.StatusOK, data)
}

func validateHero(hero Database.SuperheroQ) (bool, string) {
	err := validate.Struct(hero)
	message := ""
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			fmt.Println("Namespace:", err.Namespace())
			fmt.Println("Field:", err.Field())
			fmt.Println("StructNameSpace:", err.StructNamespace())
			fmt.Println("StructField:", err.StructField())
			fmt.Println("Tag:", err.Tag())
			fmt.Println("Actual Tag:", err.ActualTag())
			fmt.Println("Kind:", err.Kind())
			fmt.Println("Type:", err.Type())
			fmt.Println("Value:", err.Value())
			fmt.Println("Param:", err.Param())
			// message = message + string(err.Namespace()) + string(err.Field()) + string(err.StructNamespace()) + string(err.StructField()) + string(err.Tag()) + string(err.ActualTag()) + string(err.Kind()) + string(err.Param())
			fmt.Println()
			message = message + string(err.StructField()) + " " + string(err.ActualTag()) + "\n"
			// message = message + string(err.Namespace())
			// message = "there is error is validateHero()"
		}
		return false, message
	}
	return true, "no error"
}

func existanceActualName(fl validator.FieldLevel) bool {
	return Database.CheckExistActualName(fl.Field().String())
}
func existanceName(fl validator.FieldLevel) bool {
	return Database.CheckExistName(fl.Field().String())
}
