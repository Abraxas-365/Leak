package ports

import (
	"github/abraxas-365/Leak/pkg/waiting-list/domain/models"

	"github.com/google/uuid"
)

type Repository interface {
	CreateFollow(models.Follow) error
	DeleteFollow(crushId uuid.UUID, follower uuid.UUID) error //find the follow filtering with bouth uuid
	GetWaitingList(userId uuid.UUID) (models.WaitingList, error)
	GetFollow(crushId uuid.UUID, followerId uuid.UUID) (models.Follow, bool, error)
}
