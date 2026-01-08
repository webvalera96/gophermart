package service

import (
	"github.com/golang-jwt/jwt/v5"
)

const secret_salt = "secret_salt"

func IssueToken(password string) (string, error) {
	key := password + secret_salt

	t := jwt.New(jwt.SigningMethodHS256)

	s, err := t.SignedString(key)
	if err != nil {
		return "", err
	}

	return s, nil
}

func ValidateToken(password string, tokenString string) error {
	key := password + secret_salt

	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return err
	}

	return nil
}
