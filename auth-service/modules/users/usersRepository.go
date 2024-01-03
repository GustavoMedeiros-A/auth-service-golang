package users

import (
	"auth-service/modules/models"

	"gorm.io/gorm"
)

type UserConfig struct {
	DB *gorm.DB
}

func NewUserModel(db *gorm.DB) UserConfig {
	return UserConfig{
		DB: db,
	}
}

func (repo *UserConfig) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := repo.DB.Where("email = ?", email).First(&user)
	return &user, result.Error
}

func (repo *UserConfig) Create(user models.User) error {
	result := repo.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil

}
