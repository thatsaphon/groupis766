package main

import (
	"goMongo/model"
	"goMongo/mongo"
	"net/http"

	"github.com/gin-gonic/gin"
)

// "bson"

type menu struct {
	_id     string `bson:"_id"`
	user    string `bson:"user"`
	title   string `bson:"title"`
	slug    string `bson:"slug"`
	summary string `bson:"summary"`
	Type    string `bson:"type"`
	created string `bson:"created"`
	updated string `bson:"updated"`
	content string `bson:"content"`
}

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/menu", func(c *gin.Context) {
		var response model.ResponseModel
		menues := mongo.GetMenu()
		response.Data = menues

		c.JSON(http.StatusOK, response)
	})
	r.GET("/recipe", func(c *gin.Context) {
		var response model.ResponseModel
		menues := mongo.GetRecipe()
		response.Data = menues

		c.JSON(http.StatusOK, response)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
