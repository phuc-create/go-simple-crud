package controllers

import (
	"errors"
	"github.com/phuc-create/go-simple-crud/models"
)

var users = make([]*models.User, 0)

func removeSpecificElInArr(arr []*models.User, index int) []*models.User {
	return append(arr[:index], arr[index+1:]...)

}

func GetAllUsers() []*models.User {
	return users

}

func CreateUser(user *models.User) bool {
	if user.ID != "" && user.Username != "" && user.Password != "" {
		if userFind, _ := FindUser(user.ID); userFind == nil {
			users = append(users, user)
			return true
		}
	}

	return false
}

func UpdateUser(newUser *models.User) bool {
	for index, user := range users {
		if user.ID == newUser.ID {
			users[index] = newUser
			return true
		}
	}

	return false
}
func DeleteUser(id string) bool {
	for idx, user := range users {
		if user.ID == id {
			users = removeSpecificElInArr(users, idx)
			return true
		}
	}
	return false
}

func FindUser(id string) (*models.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, errors.New("user does not exist")
}
