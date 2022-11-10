package application

import "github/abraxas-365/Leak/pkg/waiting-list/domain/models"

func (a *app) FollowACrush(newFollow models.Follow) error {
	//check if follow already exist
	_, exist, err := a.repo.GetFollow(newFollow.Crush, newFollow.Follower)
	if err != nil {
		return err
	}
	if !exist {
		if err := a.repo.CreateFollow(newFollow); err != nil {
			return err
		}
	}

	return nil
}
