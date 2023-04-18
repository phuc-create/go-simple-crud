package controllers

import (
	"github.com/phuc-create/go-simple-crud/models"
	"strconv"
	"testing"
)

func TestImplement_CreateUser(t *testing.T) {
	i := implement{}
	// Case 1: Valid user info and ID not used before
	user := models.User{
		ID:       strconv.Itoa(1),
		Username: "tester",
		Password: "password",
	}
	result, err := i.CreateUser(user)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if result.ID != user.ID {
		t.Errorf("Expected user ID to be %s, but got %s", user.ID, result.ID)
	}
	if result.Username != user.Username {
		t.Errorf("Expected username to be %s, but got %s", user.Username, result.Username)
	}
	if result.Password != user.Password {
		t.Errorf("Expected password to be %s, but got %s", user.Password, result.Password)
	}
	// Case 2: Invalid user info
	user = models.User{
		ID:       strconv.Itoa(2),
		Username: "",
		Password: "password",
	}
	_, err = i.CreateUser(user)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
	// Case 3: username used before
	user = models.User{
		ID:       strconv.Itoa(3),
		Username: "tester",
		Password: "password2",
	}
	_, err = i.CreateUser(user)
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}
