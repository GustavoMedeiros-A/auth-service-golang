package common

import (
	"gorm.io/gorm"
)

type Config struct {
	DB *gorm.DB
}

func NewConfig(db *gorm.DB) *Config {
	return &Config{DB: db}
}
