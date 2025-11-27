package utils

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"school_management_api/internal/models"
	"strings"

	"golang.org/x/crypto/argon2"
)

// VerifyPassword verifies the provided password against the stored hash.
func VerifyPassword(user models.Executive, password string) error {
	// split stored password into salt and hash
	parts := strings.Split(user.Password, ".")
	if len(parts) != 2 {
		return ErrorHandler(errors.New("invalid stored password format"), "internal server error")
	}

	saltBase64, hashBase64 := parts[0], parts[1]
	salt, err := base64.StdEncoding.DecodeString(saltBase64)
	if err != nil {
		return ErrorHandler(err, "failed to decode the salt")
	}

	hashedPassword, err := base64.StdEncoding.DecodeString(hashBase64)
	if err != nil {
		return ErrorHandler(err, "failed to decode the hashed password")
	}

	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	if len(hash) != len(hashedPassword) {
		return ErrorHandler(errors.New("hashed password length mismatch"), "password verification failed")
	}

	// constant time comparison to prevent timing attacks
	if subtle.ConstantTimeCompare(hash, hashedPassword) != 1 {
		return ErrorHandler(errors.New("incorrect password"), "password verification failed")
	}
	return nil
}
