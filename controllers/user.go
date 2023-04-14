package controllers

import (
	"errors"
	"github.com/phuc-create/go-simple-crud/models"
)

var users = make([]*models.User, 0)

func removeSpecificElInArr(arr []*models.User, index int) []*models.User {
	return append(arr[:index], arr[index+1:]...)

}

func (i implement) GetAllUser() ([]*models.User, error) {
	return users, nil

}

func (i implement) CreateUser(user *models.User) (models.User, error) {
	if user.ID == "" || user.Username == "" || user.Password == "" {
		return models.User{}, errors.New("missing information. please check again")
	}

	_, err := i.GetUserByID(user.ID)
	if err != nil {
		users = append(users, user)
		return models.User{
			ID:       user.ID,
			Username: user.Username,
			Password: user.Password,
		}, nil
	}
	return models.User{}, err
}

func (i implement) GetUserByID(userID string) (models.User, error) {
	for _, user := range users {
		if user.ID == userID {
			return models.User{
				ID:       user.ID,
				Username: user.Username,
				Password: user.Password,
			}, nil
		}
	}

	return models.User{}, errors.New("user does not exist")
}

func (i implement) DeleteUser(userID string) (bool, error) {
	for idx, user := range users {

		if user.ID == userID {
			users = removeSpecificElInArr(users, idx)
			return true, nil
		}
	}
	return false, nil
}

func (i implement) UpdateUserByID(user *models.User) (models.User, error) {
	// if user not found return err, check by id
	for _, usr := range users {
		if usr.ID == user.ID {
			usr.Username = user.Username
			usr.Password = user.Password
			return models.User{
				ID:       user.ID,
				Username: user.Username,
				Password: user.Password,
			}, nil
		}
	}
	return models.User{}, errors.New("user not found")
}
