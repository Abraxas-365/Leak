package application

import (
	"github/abraxas-365/Leak/pkg/waiting-list/domain/models"

	"github.com/google/uuid"
)

func (a *app) ShowWaitingList(userId uuid.UUID) (models.WaitingList, error) {
	return a.repo.GetWaitingList(userId)
}
