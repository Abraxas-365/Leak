package application

import (
	"github/abraxas-365/Leak/internal/jwtauth"
	"github/abraxas-365/Leak/pkg/auth/domain/models"
)

func (a *app) RegisterWithInstagram(register models.RegisterWithPlatform) (string, error) {

	user, err := a.userApp.RegisterUser(register.Token, register.Password)
	if err != nil {
		return "", err
	}

	token, err := jwtauth.GereteToken(user.Id, user.InstagramToken)
	if err != nil {
		return "", err
	}

	return token, nil
}
