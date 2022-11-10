package application

import (
	"github/abraxas-365/Leak/pkg/auth/domain/models"
	user "github/abraxas-365/Leak/pkg/user/application"
)

type Application interface {
	RegisterWithInstagram(register models.RegisterWithPlatform) (string, error) //return jwt
}

type app struct {
	userApp user.Application
}

func AppFactory(userApp user.Application) Application {
	return &app{
		userApp: userApp,
	}
}
