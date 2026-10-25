package auth

import (
	"context"

	database "github.com/Mikhail-Tal63/Rooksy/database/db"
)

type AuthRepository struct {
	query *database.Queries
}

func (r *AuthRepository) NewAuthRepository(query *database.Queries) *AuthRepository {
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
