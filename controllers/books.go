package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"makkonen.com/go-mongo/config"
)

type Person struct {
	Name        string
	ReadingList []string
	Email       string
}

func GetAllBooks(c *gin.Context) {
	client := config.GetClient()
	books := client.Database("gr-clone").Collection("users")
	results := []Person{}
	cur, err := books.Find(context.Background(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		result := Person{}
		//books := []models.Book{}
		//book := models.Book{}
		err := cur.Decode(&result)
		//err := cur.Decode(&book)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result...
		results = append(results, result)
		fmt.Println(result)
		// To get the raw bson bytes use cursor.Current
		//raw := cur.Current
		// do something with raw...
	}
	/*if err := cur.Err(); err != nil {
		return err
	}*/
	c.IndentedJSON(http.StatusOK, results)
}

func GetBook(c *gin.Context) {

}

func NewBook(c *gin.Context) {

}
