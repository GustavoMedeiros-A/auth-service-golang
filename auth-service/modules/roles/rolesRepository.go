package roles

import (
	"auth-service/common"
	"auth-service/modules/models"
)

type RolesRepository struct {
	db *common.Config
}

func NewRolesRepository(config *common.Config) *RolesRepository {
	return &RolesRepository{db: config}
}

func (repo *RolesRepository) CreateRoles(Roles *models.Roles) error {
	result := repo.db.DB.Create(&Roles)
	return result.Error
}

func (repo *RolesRepository) FindAllRoles() ([]*models.Roles, error) {
	var roles []*models.Roles

	result := repo.db.DB.Preload("Permissions").Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}

	return roles, nil
}
