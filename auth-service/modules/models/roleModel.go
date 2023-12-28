package models

import "gorm.io/gorm"

type Roles struct {
	gorm.Model
	Name        string
	Permissions []Permissions `gorm:"many2many:role_permissions;"`
}
