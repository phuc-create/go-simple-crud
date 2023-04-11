package models

import (
	"errors"
	"fmt"

	"github.com/phuc-create/go-simple-crud/entities"
)

var listUser = make([]*entities.User, 0)

type ObjSimple struct {
}

func removeSpecificElInArr(arr []*entities.User, index int) []*entities.User {
	idea := ObjSimple{}
	fmt.Println(idea)
	return append(arr[:index], arr[index+1:]...)

}

func CreateUser(user *entities.User) bool {
	if user.ID != "" && user.Username != "" && user.Password != "" {
		if userFind, _ := FindUser(user.ID); userFind == nil {
			listUser = append(listUser, user)
			return true
		}
	}

	return false
}

func UpdateUser(newUser *entities.User) bool {
	for index, user := range listUser {
		if user.ID == newUser.ID {
			listUser[index] = newUser
			return true
		}
	}

	return false
}
func DeleteUser(id string) bool {
	for idx, user := range listUser {
		if user.ID == id {
			listUser = removeSpecificElInArr(listUser, idx)
			return true
		}
	}
	return false
}

func FindUser(id string) (*entities.User, error) {
	for _, user := range listUser {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, errors.New("user does not exist")
}

func GetAllUser() []*entities.User {
	return listUser
}

// Update
// Delete
// Create
// Read
