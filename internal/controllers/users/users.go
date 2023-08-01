package users

import (
	"github.com/phuc-create/go-simple-crud/models"
)

var users = make([]*models.User, 0)

func removeSpecificElInArr(arr []*models.User, index int) []*models.User {
	return append(arr[:index], arr[index+1:]...)

}

//func (i implement) CheckUserExist(username string) bool {
//	statement := "SELECT username FROM user_account WHERE username=$1"
//	err := i.db.QueryRow(statement, username).Scan(&username)
//	if err != nil {
//		return false
//	}
//	return true
//}

//func (i implement) GetUserByID(userID string) (models.User, error) {
//	var users models.User
//	statement := "SELECT * FROM user_account WHERE id=$1"
//	err := i.db.QueryRow(statement, userID).Scan(&users.ID, &users.Username, &users.Password, &users.CreatedAt, &users.UpdatedAt)
//	if err != nil {
//		if err != sql.ErrNoRows {
//			return models.User{}, err
//		}
//		return models.User{}, ErrUserDoesNotExist
//	}
//	return users, nil
//}

//func (i implement) DeleteUser(userID string) (bool, error) {
//	for idx, users := range users {
//
//		if users.ID == userID {
//			users = removeSpecificElInArr(users, idx)
//			return true, nil
//		}
//	}
//	return false, nil
//}

//func (i implement) UpdateUserByID(users *models.User) (models.User, error) {
//	// Check if users exists in the slice
//	for _, usr := range users {
//		if usr.ID == users.ID {
//			usr.Username = users.Username
//			usr.Password = users.Password
//			usr.UpdatedAt = time.Now()
//			return models.User{
//				ID:        usr.ID,
//				Username:  users.Username,
//				Password:  users.Password,
//				UpdatedAt: users.UpdatedAt,
//			}, nil
//		}
//	}
//	return models.User{}, ErrUserDoesNotExist
//}
