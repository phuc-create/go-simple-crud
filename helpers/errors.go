package helpers

import (
	"fmt"
	"net/http"
)

const (
	InternalError         = "Internal Error"
	ErrSomethingWentWrong = "Something went wrong"
)

type ErrResponse struct {
	Status  int    `json:"status,omitempty"`
	Code    string `json:"message,omitempty"`
	Message string `json:"description,omitempty"`
}

var DefaultError = &ErrResponse{
	Status:  http.StatusBadRequest,
	Code:    InternalError,
	Message: ErrSomethingWentWrong,
}

func (e ErrResponse) Error() string {
	return fmt.Sprintf("Status: [%d], Code: [%s], Desc: [%s]", e.Status, e.Message, e.Message)
}
