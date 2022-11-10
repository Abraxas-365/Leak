package application

import (
	"github.com/google/uuid"
)

func (a *app) DeleteAFollow(crushId uuid.UUID, followerId uuid.UUID) error {
	return a.repo.DeleteFollow(crushId, followerId)
}
