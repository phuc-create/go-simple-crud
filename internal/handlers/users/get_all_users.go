package users

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/phuc-create/go-simple-crud/models"
	"net/http"
)

type Data struct {
	Users []models.User `json:"users"`
}
type Response struct {
	Status int    `json:"status"`
	Data   Data   `json:"data"`
	Errors string `json:"errors"`
}

func responseJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
	return
}

func validate() {

}

// GetAllUser return all users available in DB
func (h Handler) GetAllUser(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		list, err := h.controllers.GetAllUsersController(ctx)
		if err != nil {
			responseJSON(w, http.StatusBadRequest, err)
			return
		}
		fmt.Println(list)
		responseJSON(w, http.StatusOK, Response{Status: http.StatusOK, Data: Data{Users: list}})
	}
}
