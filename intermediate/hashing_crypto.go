package main

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {

	// Hashing
	password := "password123"
	hash256 := sha256.Sum256([]byte(password))
	hash512 := sha512.Sum512([]byte(password))
	fmt.Println("SHA-256 Hashed Password:", hash256)
	fmt.Println("SHA-512 Hashed Password:", hash512)

	// Print in Hex
	fmt.Printf("SHA-256 Hash hex value: %x\n", hash256)
	fmt.Printf("SHA-512 Hash hex value: %x\n", hash512)

	fmt.Println("")
	fmt.Println("------ Salting ---------")
	fmt.Println("")

	// Salting - adds layer of security - a random byte values
	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error generating salt", err)
		return
	}

	// Hash the password
	hash_password := hashPassword(password, salt)

	// Store the salt and password into DB
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt:", saltStr)       // simulate as storing in db
	fmt.Println("Hash:", hash_password) // simulate as storing in db

	// Verify
	// Retrieve the saltStr and decode it
	decodedStr, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Error in decoding salt!", err)
		return
	}

	loginHash := hashPassword("password124", decodedStr)

	if loginHash == hash_password {
		fmt.Println("Password is correct. You are logged in.")
	} else {
		fmt.Println("Incorrect password. Please check your credentials.")
	}
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func hashPassword(password string, salt []byte) string {
	saltPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
