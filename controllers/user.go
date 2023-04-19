package controllers

import (
	"errors"
	"github.com/phuc-create/go-simple-crud/helpers"
	models "github.com/phuc-create/go-simple-crud/models"
)

var users = make([]*models.User, 0)

func removeSpecificElInArr(arr []*models.User, index int) []*models.User {
	return append(arr[:index], arr[index+1:]...)

}

func validateInfoUser(user UserInput) error {
	if user.ID == "" || user.Username == "" || user.Password == "" {
		return errors.New("missing information. please check again")
	}

	isWhiteSpace := helpers.ContainWhiteSpace(user.Password)
	if isWhiteSpace {
		return errors.New("password should not contain white space")
	}

	for _, usr := range users {
		if usr.Username == user.Username {
			return errors.New("user already exist")
		}
	}
	return nil
}

func (i implement) GetAllUser() ([]*models.User, error) {
	return users, nil

}

func (i implement) CreateUser(user models.User) (UserResponse, error) {
	if err := validateInfoUser(UserInput{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}); err != nil {
		return UserResponse{}, err
	}

	if _, err := i.GetUserByID(user.ID); err != nil {
		return UserResponse{}, err
	}

	users = append(users, &user)
	return UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}, nil
}

func (i implement) GetUserByID(userID string) (models.User, error) {
	for _, user := range users {
		if user.ID == userID {
			return models.User{
				ID:        user.ID,
				Username:  user.Username,
				Password:  user.Password,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
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
	// Check if user exists in the slice
	for _, usr := range users {
		if usr.ID == user.ID {
			usr.Username = user.Username
			usr.Password = user.Password
			return models.User{
				ID:       usr.ID,
				Username: user.Username,
				Password: user.Password,
			}, nil
		}
	}
	// Return error if user not found
	return models.User{}, errors.New("user not found")
}
