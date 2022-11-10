package application

import (
	"github/abraxas-365/Leak/pkg/waiting-list/domain/models"
	"github/abraxas-365/Leak/pkg/waiting-list/domain/ports"

	"github.com/google/uuid"
)

type Application interface {
	FollowACrush(models.Follow) error
	DeleteAFollow(crushId uuid.UUID, followerId uuid.UUID) error
	ShowWaitingList(userId uuid.UUID) (models.WaitingList, error)
}

type app struct {
	repo ports.Repository
}

func AppFactory(repo ports.Repository) Application {
	return &app{
		repo: repo,
	}
}
