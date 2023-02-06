package service

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type Top struct {
	Username string
	Amount   int
}

func (s *Service) GetTopFollowings(ctx context.Context) (string, error) {
	topFollowings, err := s.db.GetTopFollowings(context.Background(), time.Hour*24)
	if err != nil {
		return "", err
	}

	if len(topFollowings) == 0 {
		return "", errors.New("no data, try later")
	}

	prettyTop := "<b>Top follows:</b>"
	for ind, following := range topFollowings {
		user, err := s.GetUserByUsername(ctx, following.Username)
		if err != nil {
			return "", err
		}

		topFollowers, err := s.db.GetTopFollowers(ctx, user.Username, time.Hour*24)
		if err != nil {
			return "", err
		}

		prettyTop += fmt.Sprintf("\n\n<b>%d. %s (%d)</b>\nTop followers: ", ind+1, user.Username, following.Amount)
		if len(topFollowers) > 5 {
			topFollowers = topFollowers[:5]
		}
		for _, top := range topFollowers {
			prettyTop += fmt.Sprintf("<b>%s</b> ", top)
		}
	}

	return prettyTop, nil
}
