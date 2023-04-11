package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model `json:"gorm_._model"`
	ID         string `json:"id,omitempty"`
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
}
