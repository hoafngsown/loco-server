package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type Validator = func(r *http.Request) bool

var secretKey = []byte("access-token-secret-key")

func NewJWTAuthorizationMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		middleware := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bearerToken := r.Header.Get("Authorization")

			if bearerToken == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}

			bearerToken = bearerToken[len("Bearer "):]

			decodedToken, err := jwt.Parse(bearerToken, func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, http.ErrNotSupported
				}

				return secretKey, nil
			})

			fmt.Printf("decodedToken: %v", decodedToken)

			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}

			ctx := context.WithValue(r.Context(), "user", "123")

			next.ServeHTTP(w, r.WithContext(ctx))
		})

		return middleware
	}
}
