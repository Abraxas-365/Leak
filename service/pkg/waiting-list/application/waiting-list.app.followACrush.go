package application

import "github/abraxas-365/Leak/pkg/waiting-list/domain/models"

func (a *app) FollowACrush(newFollow models.Follow) error {
	//check if follow already exist
	_, exist, err := a.repo.GetFollow(newFollow.Crush, newFollow.Follower)
	if err != nil {
		return err
	}

	//check if the crush is in follower list
	follow, crushIsInFollowerList, err := a.repo.GetFollow(newFollow.Follower, newFollow.Crush)
	if err != nil {
		return err
	} else if crushIsInFollowerList {
		newFollow.IsInCrushList = true
		follow.IsInCrushList = true
		a.repo.UpdateFollow(follow)
	}

	if !exist {
		if err := a.repo.CreateFollow(newFollow); err != nil {
			return err
		}
	}

	return nil
}
