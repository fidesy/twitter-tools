package service

import (
	"context"
	"github.com/fidesy/twitter-tools/internal/models"
	"github.com/g8rswimmer/go-twitter/v2"
	"log"
)

// GetIDByUsername
func (s *Service) GetIDByUsername(ctx context.Context, username string) (string, error) {
	var user = new(models.User)
	s.db.Find(&user).Where("username = ?", username)
	if user.ID != "" {
		return user.ID, nil
	}

	resp, err := s.client.UserNameLookup(ctx, []string{username}, twitter.UserLookupOpts{})
	if err != nil {
		return "", err
	}

	user_ := resp.Raw.Users[0]
	user.ID = user_.ID
	user.Username = username
	user.Name = user_.Name
	s.db.Create(user)

	log.Printf("GetIDByUsername rate limit %d\n", resp.RateLimit.Remaining)

	return user.ID, nil
}
