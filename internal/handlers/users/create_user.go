package users

import (
	"context"
	"encoding/json"
	"github.com/phuc-create/go-simple-crud/helpers"
	"github.com/phuc-create/go-simple-crud/models"
	"net/http"
	"time"
)

type User struct {
	User models.User `json:"user"`
}
type DataResponse struct {
	Status int  `json:"status,omitempty"`
	Data   User `json:"data,omitempty"`
}

func validateInfoUser(user models.User) error {
	if user.Username == "" || user.Password == "" {
		return MissingInformation
	}
	isWhiteSpace := helpers.ContainWhiteSpace(user.Password)
	if isWhiteSpace {
		return PasswordContainWhiteSpace
	}
	return nil
}

func (h Handler) CreateUser(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		errEncode := json.NewDecoder(r.Body).Decode(&user)
		if errEncode != nil {
			responseJSON(w, http.StatusBadRequest, errorHandler(errEncode))
			return
		}

		if err := validateInfoUser(user); err != nil {
			responseJSON(w, http.StatusBadRequest, errorHandler(err))
			return
		}

		newUser, err := h.controllers.CreateUserController(ctx, models.User{
			ID:        helpers.GenerateID(),
			Username:  user.Username,
			Password:  user.Password,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
		if err != nil {
			responseJSON(w, http.StatusBadRequest, errorHandler(err))
			return
		}

		responseJSON(w, http.StatusOK, DataResponse{Status: http.StatusOK, Data: User{User: newUser}})
	}
}

//func (h Handler) CreateUser(ctx context.Context, w http.ResponseWriter, r *http.Request) {
//	var user UserInput
//	err := json.NewDecoder(r.Body).Decode(&user)
//	if err != nil {
//		helpers.ResponseWithErrs(w, http.StatusBadRequest, err.Error())
//	}
//
//	newUser, err := h.repo.Users().CreateUser(models.User{
//		ID:        helpers.GenerateID(),
//		Username:  user.Username,
//		Password:  user.Password,
//		CreatedAt: time.Now(),
//		UpdatedAt: time.Now(),
//	})
//	if err != nil {
//		helpers.ResponseWithErrs(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//	helpers.ResponseWithJSON(w, http.StatusOK, newUser)
//}
