package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Bcrypt is .
func Bcrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckBcrypt is .
func CheckBcrypt(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
