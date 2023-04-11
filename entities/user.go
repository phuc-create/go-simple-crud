package entities

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model `json:"gorm_._model,omitempty"`
	ID         string `json:"id" json:"id,omitempty"`
	Username   string `json:"username" json:"username,omitempty"`
	Password   string `json:"password" json:"password,omitempty"`
}

func (user User) ToString() string {
	return fmt.Sprintf("id:%s,\nusername:%s\npassword:%s", user.ID, user.Username, user.Password)
}
