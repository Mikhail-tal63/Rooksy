package auth

import (
	"context"

	database "github.com/Mikhail-Tal63/Rooksy/database/db"
	"github.com/Mikhail-Tal63/Rooksy/util"
	"github.com/Mikhail-Tal63/Rooksy/util/jwt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type AuthService struct {
	Repository *AuthRepository
	secret     []byte
}

func NewAuthService(repo *AuthRepository, secret []byte) *AuthService {
	return &AuthService{
		Repository: repo,
		secret:     secret,
	}
}

func (s *AuthService) Register(ctx context.Context, payload RegisterRequist) (*AuthResponce, error) {

	hashedpassword, err := util.HashPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	userID := uuid.New()
	user := &database.CreateUserParams{
		ID:           userID,
		FirstName:    payload.FirstName,
		LastName:     payload.LastName,
		Username:     payload.Username,
		Email:        payload.Email,
		PasswordHash: hashedpassword,
		ImageID:      pgtype.UUID{Valid: false},
	}

	accessToken, err := jwt.CreateJWT(s.secret, userID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.GenerateRefreshToken(s.secret, userID)
	if err != nil {
		return nil, err
	}
	newuser, err := s.Repository.CreateUser(ctx, *user)
	if err != nil {
		return nil, err
	}

	return &AuthResponce{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
		User:         *newuser,
	}, nil

}
