package auth

import database "github.com/Mikhail-Tal63/Rooksy/database/db"

type RegisterRequist struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
type LoginRequest struct {
	UsernameOrEmail string `json:"username_email"`
	Password        string `json:"password"`
}

type AuthResponce struct {
	AccessToken  string         `json:"refresh_token"`
	RefreshToken string         `json:"access_token"`
	User         *database.User `json:"user"`
}
