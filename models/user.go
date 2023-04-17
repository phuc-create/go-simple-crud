package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID       string `gorm:"ID,omitempty"`
	Username string `gorm:"Username,omitempty"`
	Password string `gorm:"Password,omitempty"`
}
