package main

import (
	"auth-service/initializers"
	"auth-service/modules/users"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB()
	initializers.LoadEnv()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()

	users.RoutesUser(router)

	router.Run()

	fmt.Println("but now that not a hello ")
}
