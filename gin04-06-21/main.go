package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"touch/Database"

	"github.com/gin-gonic/gin"
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
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	var t Database.SuperheroQ
	err := json.Unmarshal(buf[0:num], &t)
	if err != nil {
		panic(err)
	}
	if t.Name == "" {
		// fmt.Println("Need name to insert")
		c.JSON(http.StatusUnprocessableEntity, "Need name to insert")
		return
	}
	if t.Height < 0 {
		fmt.Println("age is less than 0:", t.Height)
		c.JSON(http.StatusNotFound, "age is less than 0")
		return
	}
	Database.Insert(t)
	c.JSON(http.StatusOK, reqBody)
}

func updateId(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "wrong format should be int not string")
		return
	}
	if !Database.CheckExistID(i) {
		c.JSON(http.StatusNotFound, "this id doesn't exist")
		return
	}
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	var t Database.SuperheroQ
	err = json.Unmarshal(buf[0:num], &t)
	if err != nil {
		panic(err)
	}
	t.ID = i

	if t.Height < 0 {
		c.JSON(http.StatusNotFound, "age is less than 0")
		return
	}
	Database.Update(t, i)
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
	a := Database.ViewByPage(p.Limit, p.Page)
	if a == nil {
		c.JSON(http.StatusNotFound, "this page is not available")
		return
	}
	c.JSON(http.StatusOK, a)
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
	a := Database.Search(v.Value)
	if a == nil {
		c.JSON(http.StatusOK, "No result")
		return
	}
	c.JSON(http.StatusOK, a)
}
