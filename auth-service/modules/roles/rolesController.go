package roles

import (
	"github.com/gin-gonic/gin"
)

func (rolesConfig *RolesRepository) RoutesRoles(route *gin.Engine) {
	roles := route.Group("/api")
	{
		roles.GET("/roles", rolesConfig.FindAll)
		roles.POST("/roles/create", rolesConfig.Create)
	}

}
