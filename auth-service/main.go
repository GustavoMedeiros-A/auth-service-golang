package main

import (
	"auth-service/initializers"
	"auth-service/modules/users"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	router := gin.Default()

	config := initializers.ConnectToDB()
	initializers.SyncDatabase(config)

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Link"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	usersConfig := users.NewUserRepository(config)
	usersConfig.RoutesUser(router)

	router.Run(string("0.0.0.0:8080"))

	fmt.Println("but now that not a hello ")
}
