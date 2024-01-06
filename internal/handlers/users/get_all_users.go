package users

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/phuc-create/go-simple-crud/models"
)

type Users struct {
	Users []models.User `json:"users"`
}
type Response struct {
	Status int   `json:"status"`
	Data   Users `json:"data"`
}

func responseJSONError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

func responseJSONData(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// func validate() {

// }

// GetAllUser return all users available in DB
func (h UserHandlers) GetAllUser(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		list, err := h.controllers.GetAllUsers(ctx)
		if err != nil {
			responseJSONError(w, http.StatusBadRequest, err)
			return
		}
		responseJSONData(w, Response{Status: http.StatusOK, Data: Users{Users: list}})
	}
}
