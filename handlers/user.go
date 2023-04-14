package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/phuc-create/go-simple-crud/models"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/phuc-create/go-simple-crud/helpers"
)

func GetAllUser(writer http.ResponseWriter, request *http.Request) {

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
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, err.Error())
	}
	user.ID = strconv.Itoa(rand.Intn(10000000))
	newUser, err := h.userServices.CreateUser(&user)
	if err != nil {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, err.Error())
		return
	}
	helpers.ResponseWithJSON(w, http.StatusOK, newUser)
}

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

func (h Handler) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, err.Error())
		return
	}
	if user.ID == "" {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, "could not find any user")
		return
	}
	if user.Username == "" {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, "username could not be empty")
		return
	}

	if user.Password == "" {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, "password could not be empty")
		return
	}
	result, err := h.userServices.UpdateUserByID(&user)

	if err != nil {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, err.Error())
		return
	}
	helpers.ResponseWithJSON(w, http.StatusOK, result)
}
