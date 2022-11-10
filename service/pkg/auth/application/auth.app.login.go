package application

import (
	"github/abraxas-365/Leak/internal/jwtauth"
	"github/abraxas-365/Leak/pkg/auth/domain/models"
)

func (a *app) LoginUser(login models.Login) (string, error) {
	user, err := a.userApp.LoginUser(login.Username, login.Password)
	if err != nil {
		return "", err
	}

	jwt, err := jwtauth.GereteToken(user.Id, user.InstagramToken)
	if err != nil {
		return "", err
	}

	return jwt, nil
}
