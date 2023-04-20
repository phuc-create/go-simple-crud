package user

import (
	"database/sql"
	"github.com/phuc-create/go-simple-crud/models"
	"time"
)

type UserInput struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Controllers interface {
	GetAllUser() ([]*models.User, error)
	GetUserByID(userID string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	DeleteUser(userID string) (bool, error)
	UpdateUserByID(user *models.User) (models.User, error)
}

type implement struct {
	db *sql.DB
}

// New Dependency injection
func New(db *sql.DB) Controllers {
	return implement{db: db}
}
