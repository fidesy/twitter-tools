package service

import (
	"context"
	"github.com/fidesy/twitter-tools/internal/models"
	"github.com/g8rswimmer/go-twitter/v2"
)

func (s *Service) GetFollowings(ctx context.Context, username string) ([]*models.Following, error) {
	user, err := s.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.UserFollowingLookup(ctx, user.ID, twitter.UserFollowingLookupOpts{
		MaxResults: 1000,
	})
	if err != nil {
		return nil, err
	}

	var followings = make([]*models.Following, len(resp.Raw.Users))
	for i, u := range resp.Raw.Users {
		followings[i] = &models.Following{
			UserID:            user.ID,
			Username:          username,
			FollowingID:       u.ID,
			FollowingUsername: u.UserName,
		}
		//	TODO: add user in database
	}

	return followings, nil
}
