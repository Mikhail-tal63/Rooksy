package auth

import (
	"context"
	"testing"

	database "github.com/Mikhail-Tal63/Rooksy/database/db"
	"github.com/Mikhail-Tal63/Rooksy/internal/testutil"
	"github.com/stretchr/testify/require"
)

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
