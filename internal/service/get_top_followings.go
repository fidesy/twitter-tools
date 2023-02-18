package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

type Top struct {
	Username string
	Amount   int
}

func (s *Service) GetTopFollowings(ctx context.Context) (string, error) {
	topFollowings, err := s.db.GetTopFollowings(context.Background(), time.Hour*24, 10)
	if err != nil {
		return "", err
	}

	if len(topFollowings) == 0 {
		return "", errors.New(" no data, try later")
	}

	prettyTop := "<b>Top followings:</b>"
	for ind, following := range topFollowings {
		user, err := s.GetUserByUsername(ctx, following.Username)
		if err != nil {
			log.Println("error while getting user:", following.Username, err.Error())
			continue
			//return "", err
		}

		topFollowers, err := s.db.GetTopFollowers(ctx, user.Username, time.Hour*24, 5)
		if err != nil {
			log.Println("error while getting top followers:", err.Error())
			continue
			//return "", err
		}

		prettyTop += "\n\n" + fmt.Sprintf(`<b>%d. <a href="https://twitter.com/%s">%s</a> (%d)</b>`,
			ind+1, user.Username, user.Username, following.Amount) + "\nTop followers: "

		for _, top := range topFollowers {
			prettyTop += fmt.Sprintf(`<a href="https://twitter.com/%s"><b> %s</b></a>`, top, top)
		}
	}

	return prettyTop, nil
}
