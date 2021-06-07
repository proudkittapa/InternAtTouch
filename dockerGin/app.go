package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Sp struct {
	ID          int    `bson:"ID"`
	Name        string `bson:"Name"`
	Actual_name string `bson:"Actual_name"`
	Gender      string `bson:"Gender"`
	Age         int    `bson:"Age"`
	Super_power string `bson:"Super_power"`
}

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Person struct {
	Name string
}

var name string

func main() {

	r := setupRouter()
	r.Run()

}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/insert", insert)
	r.POST("/updateById/:id", updateId)
	r.POST("/deleteById/:id", deleteId)
	r.GET("/view/:id", viewId)
	r.GET("/viewall", viewall)
	// h := CustomerHandler{}
	// h.Initialize()

	// r.GET("/customers", h.GetAllCustomer)
	// r.GET("/customers/:id", h.GetCustomer)
	// r.POST("/customers", h.SaveCustomer)
	// r.PUT("/customers/:id", h.UpdateCustomer)
	// r.DELETE("/customers/:id", h.DeleteCustomer)
	return r
}

func insert(c *gin.Context) {
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	var t Sp
	err := json.Unmarshal(buf[0:num], &t)
	if err != nil {
		panic(err)
	}
	insertDb(t)
	c.JSON(http.StatusOK, reqBody)
}

func updateId(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, i)
}

func deleteId(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		// c.Status(http.StatusNotFound)
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	c.JSON(http.StatusOK, i)
}

func viewId(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, i)
}

func viewall(c *gin.Context) {
	fmt.Println(pagination(c))
	c.JSON(http.StatusOK, "viewall")
}

func insertDb(name Sp) {
	// uri := os.Getenv("MONGODB_URI")
	uri := "mongodb://kittapa:hello@localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	collection := client.Database("test").Collection("your_collection_name")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	fmt.Println("name", name.Name)
	insertResult, err := collection.InsertOne(context.TODO(), name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	fmt.Println(name)

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
