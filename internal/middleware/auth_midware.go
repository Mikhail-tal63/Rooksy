package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Mikhail-Tal63/Rooksy/util/jso"
	"github.com/Mikhail-Tal63/Rooksy/util/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContextKey string

const UserIDKey ContextKey = "userID"

func AuthMiddleWare(secret string) func(http.Handler) http.Handler {
	sec := []byte(secret)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHader := r.Header.Get("Authorization")
			if authHader == "" {
				jso.WriteError(w, http.StatusUnauthorized, fmt.Errorf("missing authorization header"))
				return
			}
			if !strings.HasPrefix(authHader, "Bearer") {
				jso.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid authorization format"))
				return
			}
			tokenString := strings.TrimPrefix(authHader, "Bearer ")
			userID, err := jwt.VerifyJWT(sec, tokenString)
			if err != nil {
				jso.WriteError(w, http.StatusUnauthorized, err)
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
func GetUserID(ctx context.Context) (primitive.ObjectID, error) {
	userID, ok := ctx.Value(UserIDKey).(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("user id not found in context")
	}
	return userID, nil
}
