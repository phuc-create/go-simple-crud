package controllers

import (
	"errors"
	"github.com/phuc-create/go-simple-crud/models"
)

var users = make([]*models.User, 0)

func removeSpecificElInArr(arr []*models.User, index int) []*models.User {
	return append(arr[:index], arr[index+1:]...)

}

func GetAllUser() []*models.User {
	return users

}

func CreateUser(user *models.User) (models.User, error) {
	if user.ID == "" || user.Username == "" || user.Password == "" {
		return models.User{}, errors.New("missing information. please check again")
	}

	if userFind, _ := FindUser(user.ID); userFind == nil {
		users = append(users, user)
		return models.User{
			ID:       user.ID,
			Username: user.Username,
			Password: user.Password,
		}, nil
	}
	return models.User{}, errors.New("something went wrong")
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
