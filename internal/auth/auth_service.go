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

	firstname := NameValidate(payload.FirstName)
	lastname := NameValidate(payload.LastName)
	username := UsernameValidate(payload.Username)
	email := EmailValidate(payload.Email)

	if err = ValidateEmail(email); err != nil {
		return nil, err
	}
	if err = ValidateName(firstname); err != nil {
		return nil, err
	}
	if err = ValidateName(lastname); err != nil {
		return nil, err
	}
	if err = ValidateUsername(username); err != nil {
		return nil, err
	}
	if err = ValidatePassword(payload.Password); err != nil {
		return nil, err
	}

	existesUsername, err := s.Repository.GetUserByUsername(ctx, payload.Username)
	if err != nil {
		return nil, err
	}
	if existesUsername != nil {
		return nil, ErrUsernameTaken
	}
	existedEmail, err := s.Repository.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		return nil, err
	}
	if existedEmail != nil {
		return nil, ErrEmailAlreadyExists
	}

	userID := uuid.New()
	user := &database.CreateUserParams{
		ID:           userID,
		FirstName:    firstname,
		LastName:     lastname,
		Username:     username,
		Email:        email,
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
		User:         newuser,
	}, nil

}

func (s *AuthService) Login(ctx context.Context, payload LoginRequest) (*AuthResponce, error) {

	var password string
	var userID uuid.UUID
	if err := ValidateEmail(payload.UsernameOrEmail); err != nil {
		user, err := s.Repository.GetUserByUsername(ctx, payload.UsernameOrEmail)
		if err != nil {
			return nil, ErrUserNotFound
		}
		userID = user.ID
		password = user.PasswordHash
	} else {

		user, err := s.Repository.GetUserByEmail(ctx, payload.UsernameOrEmail)
		if err != nil {
			return nil, ErrUserNotFound
		}
		userID = user.ID
		password = user.PasswordHash
	}

	if !util.ComparePassword(password, []byte(payload.Password)) {
		return nil, ErrWrongPassword
	}

	accessToken, err := jwt.CreateJWT(s.secret, userID)
	if err != nil {
		return nil, err
	}
	refreshToken, err := jwt.GenerateRefreshToken(s.secret, userID)
	if err != nil {
		return nil, err
	}

	return &AuthResponce{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         nil,
	}, nil

}
