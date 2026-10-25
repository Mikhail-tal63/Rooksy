package auth

import database "github.com/Mikhail-Tal63/Rooksy/database/db"

type RegisterRequist struct {
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
}

type AuthResponce struct {
	AccessToken  string
	RefreshToken string
	User         database.User
}
