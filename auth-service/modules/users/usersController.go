package users

import (
	"auth-service/middleware"

	"github.com/gin-gonic/gin"
)

func (userConfig *UserConfig) RoutesUser(route *gin.Engine) {
	user := route.Group("/api")
	{
		user.POST("/signup", userConfig.SignUp)
		user.POST("/login", userConfig.Login)
		user.GET("/validate", middleware.RequireAuth(), Validate)
	}
}
