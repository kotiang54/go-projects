package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"school_management_api/internal/models"
	"strings"

	"golang.org/x/crypto/argon2"
)

// HashPassword hashes the password of the given executive using Argon2id.
func HashPassword(password string) (string, error) {
	if password == "" {
		return "", ErrorHandler(fmt.Errorf("password is required"), "Error inserting executive data into database")
	}

	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", ErrorHandler(errors.New("failed to generate salt"), "Error inserting executive data into database")
	}

	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	saltBase64 := base64.StdEncoding.EncodeToString(salt)
	hashBase64 := base64.StdEncoding.EncodeToString(hash)
	encodedHash := fmt.Sprintf("%s.%s", saltBase64, hashBase64)

	return encodedHash, nil
}

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
