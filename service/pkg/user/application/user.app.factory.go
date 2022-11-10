package app

import (
	"github/abraxas-365/Leak/pkg/user/domain/models"
	"github/abraxas-365/Leak/pkg/user/domain/ports"
)

type Application interface {
	RegisterUser(igToken string, password string) (models.User, error)
}

type app struct {
	userRepo ports.Repo
}

func AppFactory(repo ports.Repo) Application {
	return &app{
		userRepo: repo,
	}
}
