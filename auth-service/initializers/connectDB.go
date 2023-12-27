package initializers

import (
	"auth-service/common"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() *common.Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("cannot load .env")
	}

	dsn := os.Getenv("DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB = db

	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("connected to database")
	return common.NewConfig(db)
}
