package models

import "gorm.io/gorm"

type Permissions struct {
	gorm.Model
	Name  string
	Roles []Roles `gorm:"many2many:role_permissions;"`
}
