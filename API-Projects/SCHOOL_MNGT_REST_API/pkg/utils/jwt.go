package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// SignToken generates a JWT token for a user with specified userId, username, and role.
func SignToken(userId int, username string, role string) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	jwtExpiresIn := os.Getenv("JWT_EXPIRES_IN")

	if jwtSecret == "" {
		return "", ErrorHandler(errors.New("JWT_SECRET not set"), "Internal Error")
	}

	claims := jwt.MapClaims{
		"uid":  userId,
		"user": username,
		"role": role,
	}

	var duration time.Duration
	var err error

	if jwtExpiresIn == "" {
		duration = 15 * time.Minute
	} else {
		duration, err = time.ParseDuration(jwtExpiresIn)
		if err != nil {
			return "", ErrorHandler(err, "Internal Error")
		}
	}

	claims["exp"] = jwt.NewNumericDate(time.Now().Add(duration))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", ErrorHandler(err, "Internal Error")
	}

	return signedToken, nil
}
