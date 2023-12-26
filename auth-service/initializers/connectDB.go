package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	// dsn := "host=localhost user=postgres password=password dbname=users port=5432 sslmode=disable"
	err = godotenv.Load(".env")

	if err != nil {
		log.Fatal("cannot load .env")
	}

	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("connect to database")

}
