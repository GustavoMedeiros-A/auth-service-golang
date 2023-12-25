package initializers

import "auth-service/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
