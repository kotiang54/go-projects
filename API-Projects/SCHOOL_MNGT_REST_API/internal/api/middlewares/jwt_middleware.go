package middlewares

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"school_management_api/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
)

type ContextKey string

// JwtMiddleware validates JWT tokens in incoming requests
func JwtMiddleware(next http.Handler) http.Handler {
	fmt.Println("JWT Middleware executed")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("********* Inside JWT Middleware *********")

		token, err := r.Cookie("Bearer")
		if err != nil {
			http.Error(w, "Authorization Header Missing", http.StatusUnauthorized)
			return
		}

		jwtSecret := os.Getenv("JWT_SECRET")

		// Parse and validate the token
		parsedToken, err := jwt.Parse(token.Value, func(token *jwt.Token) (any, error) {
			// Validate the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(jwtSecret), nil
		})
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				http.Error(w, "Token has expired", http.StatusUnauthorized)
				return
			}
			utils.ErrorHandler(err, "")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Check if the token is valid
		if parsedToken.Valid {
			log.Println("Valid JWT Token")
		} else {
			http.Error(w, "Invalid Login Token", http.StatusUnauthorized)
			log.Println("Invalid Login Token:", token.Value)
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid Login Token", http.StatusUnauthorized)
		}

		ctx := context.WithValue(r.Context(), ContextKey("role"), claims["role"])
		ctx = context.WithValue(ctx, ContextKey("userid"), claims["uid"])
		ctx = context.WithValue(ctx, ContextKey("expiresAt"), claims["exp"])
		ctx = context.WithValue(ctx, ContextKey("username"), claims["user"])

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
