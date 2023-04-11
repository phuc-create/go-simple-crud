package controllers

import (
	"errors"
	"github.com/phuc-create/go-simple-crud/entities"
)

var users = make([]*entities.User, 0)

func removeSpecificElInArr(arr []*entities.User, index int) []*entities.User {
	return append(arr[:index], arr[index+1:]...)

}

func GetAllUsers() []*entities.User {
	return users

}

func CreateUser(user *entities.User) bool {
	if user.ID != "" && user.Username != "" && user.Password != "" {
		if userFind, _ := FindUser(user.ID); userFind == nil {
			users = append(users, user)
			return true
		}
	}

	return false
}

func UpdateUser(newUser *entities.User) bool {
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

func FindUser(id string) (*entities.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, errors.New("user does not exist")
}
