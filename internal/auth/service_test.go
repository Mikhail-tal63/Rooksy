package auth

import (
	"context"
	"testing"

	database "github.com/Mikhail-Tal63/Rooksy/database/db"
	"github.com/Mikhail-Tal63/Rooksy/internal/testutil"
	"github.com/stretchr/testify/require"
)

func setupAuthTest(t *testing.T) (*AuthService, *AuthRepository, *database.Queries) {
	t.Helper()

	pool := testutil.NewTestPostgres(t)
	queries := database.New(pool)
	repo := NewAuthRepository(queries)
	service := NewAuthService(repo, []byte("testing-secret"))

	return service, repo, queries
}
func TestRegister(t *testing.T) {

	pool := testutil.NewTestPostgres(t)
	query := database.New(pool)
	repo := NewAuthRepository(query)
	service := NewAuthService(repo, []byte("testing-and-shi"))

	payload := RegisterRequist{
		FirstName: "Nigga chan",
		LastName:  "ya gid3an",
		Username:  "idk",
		Email:     "a@gmail.com",
		Password:  "sweeie_fox",
	}
	ctx := context.Background()
	result, err := service.Register(ctx, payload)

	require.NoError(t, err)
	require.NotNil(t, result)

	user, err := query.GetUserByUsername(ctx, payload.Username)

	require.NoError(t, err)

	require.Equal(t, payload.Username, user.Username)
	require.Equal(t, payload.Email, user.Email)
	require.Equal(t, payload.FirstName, user.FirstName)
	require.Equal(t, payload.LastName, user.LastName)

	require.NotEmpty(t, user.PasswordHash)
	require.NotEqual(t, payload.Password, user.PasswordHash)

}

func TestRegisterWithDublcateEmails(t *testing.T) {
	ctx := context.Background()

	pool := testutil.NewTestPostgres(t)
	queries := database.New(pool)
	repo := NewAuthRepository(queries)
	service := NewAuthService(repo, []byte("hella-lowya-and-shi"))

	user1 := RegisterRequist{
		FirstName: "alzinjy",
		LastName:  "almultahib",
		Email:     "b@gmail.com",
		Username:  "smt",
		Password:  "12345678",
	}
	user2 := RegisterRequist{
		FirstName: "alzinjy",
		LastName:  "almultahib",
		Email:     "b@gmail.com",
		Username:  "smt2",
		Password:  "12345678",
	}

	_, err := service.Register(ctx, user1)

	require.NoError(t, err)
	_, err = service.Register(ctx, user2)
	require.Error(t, err)

	user, err := repo.GetUserByUsername(ctx, user1.Username)

	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, user.Email, user1.Email)

	//verify second username is not created
	_, err = repo.GetUserByUsername(ctx, user2.Username)
	require.Error(t, err)
}

func TestRegisterWithDublcateUsernames(t *testing.T) {
	ctx := context.Background()
	service, repo, _ := setupAuthTest(t)

	user1 := RegisterRequist{
		FirstName: "alzinjy",
		LastName:  "almultahib",
		Email:     "b@gmail.com",
		Username:  "smt",
		Password:  "12345678",
	}
	user2 := RegisterRequist{
		FirstName: "alzinjy",
		LastName:  "almultahib",
		Email:     "bb@gmail.com",
		Username:  "smt",
		Password:  "12345678",
	}

	_, err := service.Register(ctx, user1)
	require.NoError(t, err)
	_, err = service.Register(ctx, user2)
	require.Error(t, err)

	user, err := repo.GetUserByEmail(ctx, user1.Email)
	require.NoError(t, err)
	require.Equal(t, user.Email, user1.Email)
	require.NotNil(t, user)

	//again checking if email didnt create
	_, err = repo.GetUserByUsername(ctx, user2.Email)
	require.Error(t, err)

}
