package main

import (
	"fmt"
	"goMongo/model"
	"goMongo/mongo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/menu", func(c *gin.Context) {
		var response model.MenuResponseModel
		menues, err := mongo.GetMenu()
		if err != nil {
			fmt.Println(err)
			response.Message = "Internal Server Error"
			c.JSON(http.StatusOK, response)
		}
		response.Menus = menues
		response.Message = "Success"

		c.JSON(http.StatusOK, response)
	})

	r.POST("/menu", func(c *gin.Context) {
		var request model.CreateMenuRequest
		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = mongo.CreateMenu(request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Menu is created.",
		})
	})

	r.PUT("/menu/:title", func(c *gin.Context) {
		title := c.Param("title")
		var request model.CreateMenuRequest
		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = mongo.UpdateMenu(title, request)
		c.JSON(http.StatusOK, gin.H{
			"message": "Menu is updated.",
		})
	})

	r.DELETE("/menu/:title", func(c *gin.Context) {
		title := c.Param("title")

		err := mongo.DeleteMenu(title)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Menu is deleted.",
		})
	})

	r.GET("/recipe", func(c *gin.Context) {
		var response model.RecipeResponseModel
		menues, err := mongo.GetRecipe()
		if err != nil {
			fmt.Println(err)
			response.Message = "Internal Server Error"
			c.JSON(http.StatusOK, response)
		}
		response.Recipes = menues
		response.Message = "Success"

		c.JSON(http.StatusOK, response)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
