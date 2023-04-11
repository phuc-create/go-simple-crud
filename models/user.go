package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
