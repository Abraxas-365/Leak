package jwtauth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func GereteToken(id uuid.UUID, igToken string) (string, error) {

	secret := "JWT_SECRET_KEY"
	claims := jwt.MapClaims{
		"ig_token": igToken,
		"id":       id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}
