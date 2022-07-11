package crypt

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Encrypt string
func Encrypt(str string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func SignToken(username string) (string, error) {
	// Create the Claims
	claims := jwt.StandardClaims{
		// Expires after one day
		ExpiresAt: 86400,
		Issuer:    username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	fmt.Println(os.Getenv("TOKEN_SECRET"))

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}
