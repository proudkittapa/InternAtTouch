package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
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
	var t Database.SuperheroQ
	if err := c.ShouldBindJSON(&hero); err != nil{
		c.JSON(http.StatusBadRequest, "can't bind")
		return
	}
	
	val, message := validate(hero)
	if !val{
		c.JSON(http.StatusBadRequest, message)
		return
	}
	Database.Insert(hero)
	c.JSON(http.StatusOK, message)
}

func updateId(c *gin.Context) {
	id := c.Param("id")
	// i, err := strconv.Atoi(id)
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, "wrong format should be int not string")
	// 	return
	// }
	if !Database.CheckExistID(id) {
		c.JSON(http.StatusNotFound, "this id doesn't exist")
		return
	}
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	var hero Database.SuperheroQ
	err = json.Unmarshal(buf[0:num], &t)
	if err != nil {
		panic(err)
	}
	// hero.ID = i

	// if t.Age < 0 {
	// 	c.JSON(http.StatusNotFound, "age is less than 0")
	// 	return
	// }
	Database.Update(hero, ID)
	c.JSON(http.StatusOK, reqBody)
}

func deleteId(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "wrong format should be int not string")
		return
	}
	if !Database.CheckExistID(i) {
		// fmt.Println("this id doesn't exist")
		c.JSON(http.StatusNotFound, "this id doesn't exist")
		return
	}
	Database.Delete(i)
	c.JSON(http.StatusOK, i)
}

func viewId(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "wrong format should be int not string")
		return
	}
	if !Database.CheckExistID(i) {
		// fmt.Println("this id doesn't exist")
		c.JSON(http.StatusNotFound, "this id doesn't exist")
		return
	}

	var a Database.SuperheroQ = Database.View(i) //return struct
	c.JSON(http.StatusOK, a)
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

// func checkDate(name string) bool {
// 	myDateString := name
// 	fmt.Println("My Starting Date:\t", myDateString)
// 	fmt.Printf("%T\n", myDateString)

// 	// Parse the date string into Go's time object
// 	// The 1st param specifies the format, 2nd is our date string
// 	myDate, err := time.Parse("2006-01-02", myDateString)
// 	fmt.Printf("%T\n", myDateString)
// 	if err != nil {
// 		return false
// 	}
// 	// Format uses the same formatting style as parse, or we can use a pre-made constant
// 	fmt.Println("My Date Reformatted:\t", myDate.Format(time.RFC822))
// 	// fmt.Printf("%T\n", myDateString)
// 	// In Y-m-d
// 	fmt.Println("Just The Date:\t\t", myDate.Format("2006-01-02"))
// 	// fmt.Printf("%T\n", myDateString)
// 	return true
// }

func validate(hero Database.SuperheroQ) bool, string{
	err := validate.Struct(hero)
	message:= ""
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
