package user

import (
	"database/sql"
	"github.com/phuc-create/go-simple-crud/helpers"
	"github.com/phuc-create/go-simple-crud/models"
	"log"
)

var users = make([]*models.User, 0)

func removeSpecificElInArr(arr []*models.User, index int) []*models.User {
	return append(arr[:index], arr[index+1:]...)

}

func validateInfoUser(user models.User) error {
	if user.ID == "" || user.Username == "" || user.Password == "" {
		return ErrMissingInformation
	}

	isWhiteSpace := helpers.ContainWhiteSpace(user.Password)
	if isWhiteSpace {
		return ErrPasswordContainWhiteSpace
	}
	return nil
}

func (i implement) GetAllUser() ([]models.User, error) {
	var users []models.User
	statement := "SELECT * FROM user_account"
	rows, err := i.db.Query(statement)
	if err != nil {
		return []models.User{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

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
		users = append(users, user)
	}
	return users, nil
}

func (i implement) CheckUserExist(username string) bool {
	statement := "SELECT username FROM user_account WHERE username=$1"
	err := i.db.QueryRow(statement, username).Scan(&username)
	if err != nil {
		return false
	}
	return true
}

func (i implement) CreateUser(user models.User) (models.User, error) {
	if err := validateInfoUser(user); err != nil {
		return models.User{}, err
	}
	if exist := i.CheckUserExist(user.Username); exist {
		return models.User{}, ErrUserAlreadyExist
	}
	statement := "INSERT INTO user_account (id,username,password,created_at,updated_at) VALUES ($1,$2,$3,$4,$5)"
	if _, err := i.db.Exec(
		statement,
		user.ID,
		user.Username,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	); err != nil {
		return models.User{}, err
	}
	return user, nil
}

//func (i implement) GetUserByID(userID string) (models.User, error) {
//	var user models.User
//	statement := "SELECT * FROM user_account WHERE id=$1"
//	err := i.db.QueryRow(statement, userID).Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
//	if err != nil {
//		if err != sql.ErrNoRows {
//			return models.User{}, err
//		}
//		return models.User{}, ErrUserDoesNotExist
//	}
//	return user, nil
//}

//func (i implement) DeleteUser(userID string) (bool, error) {
//	for idx, user := range users {
//
//		if user.ID == userID {
//			users = removeSpecificElInArr(users, idx)
//			return true, nil
//		}
//	}
//	return false, nil
//}

//func (i implement) UpdateUserByID(user *models.User) (models.User, error) {
//	// Check if user exists in the slice
//	for _, usr := range users {
//		if usr.ID == user.ID {
//			usr.Username = user.Username
//			usr.Password = user.Password
//			usr.UpdatedAt = time.Now()
//			return models.User{
//				ID:        usr.ID,
//				Username:  user.Username,
//				Password:  user.Password,
//				UpdatedAt: user.UpdatedAt,
//			}, nil
//		}
//	}
//	return models.User{}, ErrUserDoesNotExist
//}
