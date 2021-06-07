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
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
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
	r.POST("/insert", insert)
	r.POST("/update/:id", updateId)
	r.POST("/delete/:id", deleteId)
	r.GET("/view/:id", viewId)
	r.GET("/viewall", viewall)
	return r
}

func insert(c *gin.Context) {
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	var t Database.Superhero_q
	err := json.Unmarshal(buf[0:num], &t)
	if err != nil {
		panic(err)
	}
	if t.Name == "" {
		fmt.Println("Need name to insert")
		c.Status(http.StatusNotFound)
		return
	}
	if t.Age < 0 {
		fmt.Println("age is less than 0:", t.Age)
		c.Status(http.StatusNotFound)
		return
	}
	Database.Insert(t)
	c.JSON(http.StatusOK, reqBody)
}

func updateId(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	//check if id exist
	if !Database.Check_exist_ID(i) {
		fmt.Println("this id doesn't exist")
		c.Status(http.StatusNotFound)
		return
	}
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	var t Database.Superhero_q
	err = json.Unmarshal(buf[0:num], &t)
	if err != nil {
		panic(err)
	}
	t.ID = i
	fmt.Println("Name:", t.Name == "")
	fmt.Println("Age:", t.Age == 0)
	fmt.Println("Gender:", t.Gender == "")
	if t.Age < 0 {
		fmt.Println("age is less than 0:", t.Age)
		c.Status(http.StatusNotFound)
		return
	}
	Database.Update(t, i)
	c.JSON(http.StatusOK, reqBody)
}

func deleteId(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		// c.Status(http.StatusNotFound)
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	Database.Delete(i)
	c.JSON(http.StatusOK, i)
}

func viewId(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	var a Database.Superhero_q = Database.View(i) //return struct
	c.JSON(http.StatusOK, a)
}

func viewall(c *gin.Context) {
	p := pagination(c)
	a := Database.Viewall(p.Limit, p.Offset)
	c.JSON(http.StatusOK, a)
}

func pagination(c *gin.Context) Pagination {
	limit := 2
	page := 1
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		}
	}
	return Pagination{
		Limit:  limit,
		Offset: page,
	}
}
