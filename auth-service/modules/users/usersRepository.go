package users

import (
	"auth-service/common"
	"auth-service/modules/models"
)

type Config struct {
	db *common.Config
}

func NewUserRepository(config *common.Config) *Config {
	return &Config{db: config}
}

func (repo *Config) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := repo.db.DB.Where("email = ?", email).First(&user)
	return &user, result.Error
}

func (repo *Config) Create(user models.User) error {
	result := repo.db.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil

}
