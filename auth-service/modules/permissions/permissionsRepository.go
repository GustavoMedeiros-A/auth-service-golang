package permissions

import (
	"auth-service/common"
	"auth-service/modules/models"
)

type PermissionRepository struct {
	db *common.Config
}

func NewPermissionRepository(config *common.Config) *PermissionRepository {
	return &PermissionRepository{db: config}
}

func (repo *PermissionRepository) CreatePermission(permission *models.Permissions) error {
	result := repo.db.DB.Create(&permission)
	return result.Error
}

func (repo *PermissionRepository) FindAllPermissions() ([]*models.Permissions, error) {
	var permissions []*models.Permissions

	result := repo.db.DB.Find(&permissions)
	if result.Error != nil {
		return nil, result.Error
	}

	return permissions, nil
}

func (repo *PermissionRepository) FindByIdPermission(id uint) (*models.Permissions, error) {
	var permission models.Permissions

	// Assuming your permissions have an ID field that's a uint
	result := repo.db.DB.First(&permission, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &permission, nil
}
