package service

import (
	"context"
	"log"
)

func (s *Service) GetNewFollowings(ctx context.Context, username string) ([]string, error) {
	currentFollowings, err := s.GetFollowings(ctx, username)
	if err != nil {
		return nil, err
	}

	dbFollowings, err := s.db.GetFollowingsByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if len(dbFollowings) == 0 {
		for _, following := range currentFollowings {
			err = s.db.AddFollowing(context.Background(), following)
			if err != nil {
				log.Printf("error adding following: %s", err.Error())
			}
		}
		return []string{}, nil
	}

	// fill in hashmap with all followings from the database
	var followings = make(map[string]bool)
	for _, following := range dbFollowings {
		followings[following.FollowingUsername] = true
	}

	var newFollowings []string
	for _, following := range currentFollowings {
		if _, ok := followings[following.FollowingUsername]; !ok {
			newFollowings = append(newFollowings, following.FollowingUsername)
			err = s.db.AddFollowing(context.Background(), following)
			if err != nil {
				log.Printf("error adding following: %s", err.Error())
			}
		}
	}

	return newFollowings, nil
}
