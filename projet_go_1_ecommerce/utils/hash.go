package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// hashing pw func

func HashPassword(password string) (string, error) {

	byte, err := bcrypt.GenerateFromPassword([]byte(password))

	return string(byte), err
}
func CheckPasswordHash(password, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
