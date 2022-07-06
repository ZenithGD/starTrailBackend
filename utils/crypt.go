package crypt

import (
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
