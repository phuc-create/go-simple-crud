package handlers

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/phuc-create/go-simple-crud/helpers"
	"github.com/phuc-create/go-simple-crud/models"
	"net/http"
	"time"
)

func validateInfoUser(user UserInput) error {
	if user.ID == "" {
		return errors.New("could not find any user")
	}

	if user.Username == "" {
		return errors.New("username could not be empty")
	}

	if user.Password == "" {
		return errors.New("password could not be empty")
	}
	return nil
}

// GetUserByID return user following id
func (h Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	if userID == "" {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, "User Not Found!Pls check again.")
		return
	}
	user, err := h.userServices.GetUserByID(userID)
	if err != nil {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, err.Error())
		return
	}
	helpers.ResponseWithJSON(w, http.StatusOK, user)
}

// GetAllUser return all users available in DB
func (h Handler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list, _ := h.userServices.GetAllUser()
	if len(list) < 1 {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, "Could not find any user!")
	}
	helpers.ResponseWithJSON(w, http.StatusOK, list)
}

// CreateUser create new user
func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user UserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, err.Error())
	}

	newUser, err := h.userServices.CreateUser(models.User{
		ID:        helpers.GenerateID(),
		Username:  user.Username,
		Password:  user.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, err.Error())
		return
	}
	helpers.ResponseWithJSON(w, http.StatusOK, newUser)
}

// DeleteUser delete a user
func (h Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	if userID == "" {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, "User Not Found!Pls check again.")
		return
	}
	result, err := h.userServices.DeleteUser(userID)

	if err != nil {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, err.Error())
		return
	}
	helpers.ResponseWithJSON(w, http.StatusOK, result)
}

// UpdateUserByID update user following ID
func (h Handler) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	var user UserInput
	userID := chi.URLParam(r, "userID")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, err.Error())
		return
	}
	err = validateInfoUser(UserInput{
		ID:       userID,
		Username: user.Username,
		Password: user.Password,
	})
	if err != nil {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, err.Error())
		return
	}
	result, err := h.userServices.UpdateUserByID(&models.User{
		ID:       userID,
		Username: user.Username,
		Password: user.Password,
	})

	if err != nil {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, err.Error())
		return
	}
	helpers.ResponseWithJSON(w, http.StatusOK, result)
}
