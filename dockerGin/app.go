package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
	r.POST("/updateByName/:name", updateName)
	r.POST("/deleteByName/:name", deleteName)
	r.GET("/viewall", viewall)
	// h := CustomerHandler{}
	// h.Initialize()

	// r.GET("/customers", h.GetAllCustomer)
	// // r.GET("/customers/:id", h.GetCustomer)
	// r.POST("/customers", h.SaveCustomer)
	// r.PUT("/customers/:id", h.UpdateCustomer)
	// r.DELETE("/customers/:id", h.DeleteCustomer)

	return r
}

func insert(c *gin.Context) {
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	var t Person
	err := json.Unmarshal(buf[0:num], &t)
	if err != nil {
		panic(err)
	}
	insertDb(t)
	c.JSON(http.StatusOK, reqBody)
}

func updateId(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, id)
}

func deleteId(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, id)
}

func viewall(c *gin.Context) {

	c.JSON(http.StatusOK, "viewall")
}

func updateName(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, name)
}

func deleteName(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, name)
}

func insertDb(name Person) {
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

}
