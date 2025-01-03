package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	UserName string `gorm:"unique"`
	PW       string
}
