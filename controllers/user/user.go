package user

import (
	"github.com/phuc-create/go-simple-crud/helpers"
	models "github.com/phuc-create/go-simple-crud/models"
	"log"
	"time"
)

var users = make([]*models.User, 0)

func removeSpecificElInArr(arr []*models.User, index int) []*models.User {
	return append(arr[:index], arr[index+1:]...)

}

func validateInfoUser(user UserInput) error {
	if user.ID == "" || user.Username == "" || user.Password == "" {
		return ErrMissingInformation
	}

	isWhiteSpace := helpers.ContainWhiteSpace(user.Password)
	if isWhiteSpace {
		return ErrPasswordContainWhiteSpace
	}

	for _, usr := range users {
		if usr.Username == user.Username {
			return ErrUserAlreadyExist
		}
	}
	return nil
}

func (i implement) GetAllUser() ([]*models.User, error) {
	var users []*models.User
	statement := "select * from user_account"
	rows, err := i.db.Query(statement)
	if err != nil {
		return []*models.User{}, err
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, &user)
	}
	return users, nil
}

func (i implement) CreateUser(user models.User) (models.User, error) {
	if err := validateInfoUser(UserInput{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}); err != nil {
		return models.User{}, err
	}

	users = append(users, &user)
	return models.User{
		ID:        user.ID,
		Username:  user.Username,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
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

	return models.User{}, ErrUserDoesNotExist
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
			usr.UpdatedAt = time.Now()
			return models.User{
				ID:        usr.ID,
				Username:  user.Username,
				Password:  user.Password,
				UpdatedAt: user.UpdatedAt,
			}, nil
		}
	}
	// Return error if user not found
	return models.User{}, ErrUserDoesNotExist
}
