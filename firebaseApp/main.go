package main

import (
	"firebaseApp/firebase"
	"firebaseApp/middleware"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	firebaseApp, err := firebase.InitFirebase()
	if err != nil {
		fmt.Println(err)
	}
	firebase.InitClientAuth(firebaseApp)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/user", func(c *gin.Context) {
		type RequestModel struct {
			Email       string `json:"email"`
			PhoneNumber string `json:"phoneNumber"`
			Password    string `json:"password"`
			DisplayName string `json:"displayName"`
			PhotoURL    string `json:"photoURL"`
		}
		var request RequestModel
		c.ShouldBindJSON(&request)
		fmt.Printf("%#v\n", request)
		err := firebase.CreateUser(
			request.Email,
			request.PhoneNumber,
			request.Password,
			request.DisplayName,
			request.PhotoURL,
		)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Cannot create user",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "user was created.",
		})
	})
	r.POST("/getCustomToken", func(c *gin.Context) {
		type RequestModel struct {
			Email string `json:"email"`
		}
		var request RequestModel
		c.ShouldBindJSON(&request)
		fmt.Println(request)
		user, err := firebase.GetUserByEmail(request.Email)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{
				"message": "error: cannot find email",
			})
			return
		}
		token, err := firebase.CreateCustomToken(user.UID)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{
				"message": "error: cannot create custom token",
			})
			return
		}
		c.JSON(200, gin.H{
			"token": token,
		})
	})

	r.Use(middleware.AuthMiddleware)

	//service หลัก

	r.Run()
}
