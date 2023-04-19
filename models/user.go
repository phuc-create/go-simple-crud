package models

import (
	"time"
)

type User struct {
	ID        string    `gorm:"id,primaryKey,omitempty"`
	Username  string    `gorm:"username,omitempty"`
	Password  string    `gorm:"password,omitempty"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}
