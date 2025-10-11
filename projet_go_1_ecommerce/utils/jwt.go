package utils

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JWTSecret string

func GenerateJWTSecret() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {

		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

// build a secret key
func BuildJWTSecret() error {
	var err error
	JWTSecret, err = GenerateJWTSecret()
	if err != nil {
		return err
	}
	return nil
}

// generate a jwt token out of the email and hashed password

func GenerateJWT(

}
