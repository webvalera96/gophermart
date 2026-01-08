package service

import (
	"github.com/golang-jwt/jwt/v5"
)

// TODO: generate secret key securely
const secret = "my_secret_key"

func IssueToken(login string) (string, error) {
	key := secret

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
	})

	s, err := t.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return s, nil
}

func ValidateToken(tokenString string) (string, error) {
	key := secret

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	login := token.Claims.(jwt.MapClaims)["login"].(string)

	if err != nil {
		return "", err
	}

	return login, nil
}
