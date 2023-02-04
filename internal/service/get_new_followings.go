package service

import (
	"context"
	"github.com/fidesy/twitter-tools/internal/models"
)

func (s *Service) GetFollowingActions(ctx context.Context, username string) ([]string, []string, error) {
	var dbFollowings = make([]*models.Following, 0)
	s.db.Find(&dbFollowings, models.Following{Username: username})

	currentFollowings, err := s.GetFollowings(ctx, username)
	if err != nil {
		return nil, nil, err
	}

	s.db.Unscoped().Where("username = ?", username).Delete(&models.Following{})
	for _, following := range currentFollowings {
		s.db.Create(following)
	}

	if len(dbFollowings) == 0 {
		return nil, nil, nil
	}

	// fill in map with an old followings
	var oldFollowings = make(map[string]bool, len(dbFollowings))
	for _, following := range dbFollowings {
		oldFollowings[following.FollowingID] = true
	}

	// fill in map with a current followings
	var currentFollowings_ = make(map[string]bool, len(currentFollowings))
	for _, following := range currentFollowings {
		currentFollowings_[following.FollowingID] = true
	}

	// find new followings
	var newFollowings = make([]string, 0)
	for _, following := range currentFollowings {
		if _, ok := oldFollowings[following.FollowingID]; !ok {
			newFollowings = append(newFollowings, following.FollowingUsername)
		}
	}

	// find deleted followings
	var deletedFollowings = make([]string, 0)
	for _, following := range dbFollowings {
		if _, ok := currentFollowings_[following.FollowingID]; !ok {
			deletedFollowings = append(deletedFollowings, following.FollowingUsername)
		}
	}

	return newFollowings, deletedFollowings, nil
}
