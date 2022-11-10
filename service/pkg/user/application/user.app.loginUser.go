package application

import (
	"errors"
	"github/abraxas-365/Leak/internal/leakerrs"
	"github/abraxas-365/Leak/pkg/user/domain/models"
)

func (a *app) LoginUser(username string, password string) (models.User, error) {
	//check if user exist
	user, exist, err := a.userRepo.GetUser("username", username)
	if err != nil {
		return models.User{}, err
	}
	if !exist {
		return models.User{}, errors.New(leakerrs.DocumentNotFound)
	}

	if user.Password != password {
		return models.User{}, errors.New(leakerrs.DocumentNotFound)
	}

	return user, nil
}
