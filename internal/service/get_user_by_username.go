package service

import (
	"context"
	"errors"
	"github.com/fidesy/twitter-tools/internal/models"
	"github.com/g8rswimmer/go-twitter/v2"
)

func (s *Service) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	user, err := s.db.GetUserByUsername(ctx, username)
	if err == nil {
		return user, nil
	}

	resp, err := s.client.UserNameLookup(ctx, []string{username}, twitter.UserLookupOpts{
		UserFields: []twitter.UserField{
			twitter.UserFieldDescription,
			twitter.UserFieldPublicMetrics,
			twitter.UserFieldProfileImageURL,
		},
	})
	if err != nil {
		return nil, err
	}

	if resp.Raw.Users[0] == nil {
		return nil, errors.New("account not found")
	}

	u := resp.Raw.Users[0]
	user = &models.User{
		ID:              u.ID,
		Username:        u.UserName,
		Name:            u.Name,
		Description:     u.Description,
		ProfileImageURL: u.ProfileImageURL,
		Following:       u.PublicMetrics.Following,
		Followers:       u.PublicMetrics.Followers,
		Tweets:          u.PublicMetrics.Tweets,
		IsTracked:       false,
	}
	s.db.AddUser(ctx, user)

	return user, nil
}
