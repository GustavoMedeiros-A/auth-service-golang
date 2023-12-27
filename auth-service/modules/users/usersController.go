package users

import (
	"github.com/gin-gonic/gin"
)

func (repo *Config) RoutesUser(route *gin.Engine) {
	user := route.Group("/api")
	{
		user.POST("/signup", repo.SignUp)
		user.POST("/login", repo.Login)
		user.GET("/validate", middleware.RequireAuth, Validate)
	}
}
