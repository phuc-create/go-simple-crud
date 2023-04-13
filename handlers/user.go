package handlers

import (
	"encoding/json"
	"github.com/phuc-create/go-simple-crud/models"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/phuc-create/go-simple-crud/helpers"
)

func GetAllUser(writer http.ResponseWriter, request *http.Request) {

}

func (h Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, "User Not Found!Pls check again.")
		return
	}
	user, err := h.userControllers.GetUserByID(ids[0])
	if err != nil {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, err.Error())
	}
	helpers.ResponseWithJSON(w, http.StatusOK, user)
}

func (h Handler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list, _ := h.userControllers.GetAllUser()
	if len(list) < 1 {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, "Could not find any user!")
	}
	helpers.ResponseWithJSON(w, http.StatusOK, list)
}

func (h Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, err.Error())
	}
	user.ID = strconv.Itoa(rand.Intn(10000000))
	newUser, err := h.userControllers.CreateUser(&user)
	if err != nil {
		helpers.ResponseWithErrs(w, http.StatusBadRequest, err.Error())
		return
	}
	helpers.ResponseWithJSON(w, http.StatusOK, newUser)

}
