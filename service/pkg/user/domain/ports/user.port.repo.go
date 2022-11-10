package ports

import "github/abraxas-365/Leak/pkg/user/domain/models"

type Repo interface {
	CreateUser(models.User) error
	UpdateUser(models.User) error
	GetUser(key string, value interface{}) (models.User, bool, error)
}
