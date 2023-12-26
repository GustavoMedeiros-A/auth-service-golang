package users

import (
	"auth-service/middleware"

	"github.com/gin-gonic/gin"
)

func RoutesUser(route *gin.Engine) {
	user := route.Group("/api")
	{
		user.POST("/signup", SignUp)
		user.POST("/login", Login)
		user.GET("/validate", middleware.RequireAuth, Validate)
	}
}
