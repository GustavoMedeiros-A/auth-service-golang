package roles

import (
	"auth-service/modules/models"
	"auth-service/modules/permissions"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type PermissionRepository struct {
// 	permissionRepository *permissions.PermissionRepository
// }

func (repositoryRoles *RolesRepository) Create(c *gin.Context) {

	var body struct {
		Name          string `json:"name"`
		PermissionIDs []uint `json:"permission_ids"`
	}

	// Parse JSON request body into the 'body' variable
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var permissionModel []models.Permissions
	permRepo := permissions.NewPermissionRepository(repositoryRoles.db)

	for _, id := range body.PermissionIDs {
		permission, err := permRepo.FindByIdPermission(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Permission with ID %d not found", id)})
			return
		}
		permissionModel = append(permissionModel, *permission)
	}

	roles := models.Roles{
		Name:        body.Name,
		Permissions: permissionModel,
	}

	err := repositoryRoles.CreateRoles(&roles)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	// Respond with a success message or the created Permission details
	c.JSON(http.StatusOK, gin.H{
		"message": "Role created successfully",
		"role":    roles,
	})

}

func (repo *RolesRepository) FindAll(c *gin.Context) {

	roles, err := repo.FindAllRoles()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"roles": roles})

}
