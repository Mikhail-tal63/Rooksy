package auth

import (
	"context"

	database "github.com/Mikhail-Tal63/Rooksy/database/db"
)

type AuthRepository struct {
	query *database.Queries
}

func NewAuthRepository(query *database.Queries) *AuthRepository {
	return &AuthRepository{
		query: query,
	}
}

func (r *AuthRepository) CreateUser(ctx context.Context, user database.CreateUserParams) (*database.User, error) {
	new, err := r.query.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &new, nil
}

func (r *AuthRepository) GetUserByEmail(ctx context.Context, email string) (*database.GetUserByEmailRow, error) {
	new, err := r.query.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &new, nil
}
func (r *AuthRepository) GetUserByUsername(ctx context.Context, username string) (*database.GetUserByUsernameRow, error) {
	new, err := r.query.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return &new, nil
}
