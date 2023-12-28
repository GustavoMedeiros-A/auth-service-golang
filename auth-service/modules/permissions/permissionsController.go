package permissions

import "github.com/gin-gonic/gin"

func (permissionConfig *PermissionRepository) RoutesPermission(route *gin.Engine) {
	permission := route.Group("/api")
	{
		permission.GET("/permission", permissionConfig.FindAll)
		permission.POST("/permission/create", permissionConfig.Create)
	}

}
