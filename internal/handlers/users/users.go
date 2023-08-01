package users

import (
	"time"
)

type UserInput struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetUserByID return users following id
//func (h Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
//	userID := chi.URLParam(r, "userID")
//	if userID == "" {
//		helpers.ResponseWithErrs(w, http.StatusBadRequest, "Invalid users ID!Pls check again.")
//		return
//	}
//	users, err := h.userServices.GetUserByID(userID)
//	if err != nil {
//		helpers.ResponseWithErrs(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//	helpers.ResponseWithJSON(w, http.StatusOK, users)
//}

// CreateUser create new users
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

// DeleteUser delete a users
//func (h Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
//	userID := chi.URLParam(r, "userID")
//	if userID == "" {
//		helpers.ResponseWithErrs(w, http.StatusBadRequest, "User Not Found!Pls check again.")
//		return
//	}
//	result, err := h.userServices.DeleteUser(userID)
//
//	if err != nil {
//		helpers.ResponseWithErrs(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//	helpers.ResponseWithJSON(w, http.StatusOK, result)
//}

// UpdateUserByID update users following ID
//func (h Handler) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
//	var users UserInput
//	userID := chi.URLParam(r, "userID")
//	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
//		helpers.ResponseWithErrs(w, http.StatusBadRequest, err.Error())
//		return
//	}
//	if err := validateInfoUser(UserInput{
//		ID:       userID,
//		Username: users.Username,
//		Password: users.Password,
//	}); err != nil {
//		helpers.ResponseWithErrs(w, http.StatusBadRequest, err.Error())
//		return
//	}
//	result, err := h.userServices.UpdateUserByID(&models.User{
//		ID:       userID,
//		Username: users.Username,
//		Password: users.Password,
//	})
//
//	if err != nil {
//		helpers.ResponseWithErrs(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//	helpers.ResponseWithJSON(w, http.StatusOK, result)
//}
