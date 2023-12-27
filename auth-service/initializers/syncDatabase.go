package initializers

import (
	"auth-service/common"
	"auth-service/modules/models"
)

func SyncDatabase(config *common.Config) {
	config.DB.AutoMigrate(&models.User{})
}
