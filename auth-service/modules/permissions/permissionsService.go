package permissions

import (
	"auth-service/modules/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (repo *PermissionRepository) Create(c *gin.Context) {

	var body struct {
		Name string
	}

	// Parse JSON request body into the 'body' variable
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Print(body, "this is the body")

	permission := &models.Permissions{Name: body.Name}

	err := repo.CreatePermission(permission)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	// Respond with a success message or the created Permission details
	c.JSON(http.StatusOK, gin.H{
		"message":    "Permission created successfully",
		"permission": permission,
	})

}

func (repo *PermissionRepository) FindAll(c *gin.Context) {

	permission, err := repo.FindAllPermissions()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"permission": permission})

}
