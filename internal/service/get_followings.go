package service

import (
	"context"
	"github.com/fidesy/twitter-tools/internal/models"
	"github.com/g8rswimmer/go-twitter/v2"
	"log"
)

func (s *Service) GetFollowings(ctx context.Context, username string) ([]*models.Following, error) {
	userID, err := s.GetIDByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.UserFollowingLookup(ctx, userID, twitter.UserFollowingLookupOpts{})
	if err != nil {
		return nil, err
	}

	log.Printf("GetFollowings rate limit %d\n", resp.RateLimit.Remaining)

	var followings = make([]*models.Following, len(resp.Raw.Users))
	for i, user := range resp.Raw.Users {
		followings[i] = &models.Following{
			UserID:            userID,
			Username:          username,
			FollowingID:       user.ID,
			FollowingUsername: user.UserName,
		}
		//	TODO: add user in database
	}

	return followings, nil
}
