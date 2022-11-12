package application

import (
	"errors"
	"github/abraxas-365/Leak/internal/leakerrs"
	"github/abraxas-365/Leak/pkg/user/domain/models"
)

/*
The front end will resice the token and send it to the back
in the front end call the IG api with the token to show the username, and let the user put a password
*/
func (a *app) RegisterUser(igToken string, password string) (models.User, error) {
	//create a new user
	newUser := models.NewUser(igToken, password)
	//get username from ig
	if err := newUser.GetUsername(); err != nil {
		return models.User{}, err
	}

	//check if user exist
	user, exist, err := a.userRepo.GetUser("username", newUser.UserName)
	if err != nil {
		return models.User{}, err
	}

	if !exist {
		//If user dont exist create it
		if err := a.userRepo.CreateUser(newUser); err != nil {
			return models.User{}, err
		}
		return newUser, nil
	}

	//if user exist and isTaken return user exist
	if user.IsTaken {
		return models.User{}, errors.New(leakerrs.DocumentExist)
	}
	//if user exist and is not taken update to taken, updatePassword and return it
	user.IsTaken = true
	user.InstagramToken.Token = igToken
	user.Password = password
	if err := a.userRepo.UpdateUser(user); err != nil {
		return models.User{}, err
	}

	return user, nil

}
