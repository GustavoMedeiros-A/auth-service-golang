package main

import (
	"auth-service/initializers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB()
	initializers.LoadEnv()
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ping",
		})
	})

	router.Run()

	fmt.Println("but now that not a hello ")
}
