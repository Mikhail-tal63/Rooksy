package auth

import "net/http"

type AuthErrors struct {
	msg    string
	status int
	code   string
}

func (a *AuthErrors) Error() string   { return a.msg }
func (a *AuthErrors) Code() string    { return a.code }
func (a *AuthErrors) HTTPStatus() int { return a.status }

var (
	ErrUserNotFound       error = &AuthErrors{"user not found", http.StatusNotFound, "USER_NOT_FOUND"}
	ErrEmailAlreadyExists error = &AuthErrors{"email already exists", http.StatusConflict, "EMAIL_ALREADY_EXISTS"}
	ErrUsernameTaken      error = &AuthErrors{"username already taken", http.StatusConflict, "USERNAME_TAKEN"}
	ErrInvalidCredentials error = &AuthErrors{"invalid email or password", http.StatusUnauthorized, "INVALID_CREDENTIALS"}
	ErrUnauthorized       error = &AuthErrors{"unauthorized", http.StatusUnauthorized, "UNAUTHORIZED"}
	ErrInvalidToken       error = &AuthErrors{"invalid token", http.StatusUnauthorized, "INVALID_TOKEN"}
	ErrTokenExpired       error = &AuthErrors{"token expired", http.StatusUnauthorized, "TOKEN_EXPIRED"}
	ErrWrongPassword      error = &AuthErrors{"wrong email or password", http.StatusUnauthorized, "WRONG_PASSWORD"}
)
