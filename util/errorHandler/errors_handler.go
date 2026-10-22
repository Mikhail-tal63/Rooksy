package errorhandler

import (
	"errors"
	"net/http"

	"github.com/Mikhail-Tal63/Rooksy/util/jso"
)

type AppError interface {
	error
	HTTPStatus() int
	Code() string
}

type ErrorResponse struct {
	Success bool        `json:"success"`
	Error   ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Handle(w http.ResponseWriter, err error) {
	var app AppError
	if errors.As(err, &app) {
		Write(w, app.HTTPStatus(), app.Code(), app.Error())
		return
	}
	Write(w, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "internal server error")
}

func Write(w http.ResponseWriter, status int, code string, message string) {
	_ = jso.WriteJSON(w, status, ErrorResponse{
		Success: false,
		Error: ErrorDetail{
			Code:    code,
			Message: message,
		},
	})
}
