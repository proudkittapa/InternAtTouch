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

func main() {

	r := setupRouter()
	Database.InitDB()
	r.Run()

}

func setupRouter() *gin.Engine {
	r := gin.Default()
	// r.POST("/insert", insert)
	// r.PUT("/update/:id", updateId)
	// r.DELETE("/delete/:id", deleteId)
	// r.GET("/view/:id", viewId)
	// r.GET("/viewall", viewall)
	// r.GET("/search", search)
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

	val, message := validateHero(hero)
	if !val {
		c.JSON(http.StatusBadRequest, message)
		return
	}
	Database.Insert(hero)
	c.JSON(http.StatusOK, message)
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
	c.JSON(http.StatusOK, message)
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
	data := Database.SearchName(v.Value)
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
			message = message + err.Namespace() + err.Field() + err.StructNamespace() + err.StructField() + err.Tag() + err.ActualTag() + err.Kind() + err.Type() + err.Value() + err.Param()
			fmt.Println()

		}
		return false, message
	}
	return true, "no error"
}
