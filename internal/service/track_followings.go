package service

import (
	"context"
	"fmt"
	"github.com/fidesy/twitter-tools/internal/models"
	"log"
	"time"
)

func (s *Service) TrackFollowings(actions chan<- string, usernames []string) {
	for _ = range time.Tick(time.Minute) {
		username := usernames[0]
		usernames = append(usernames[1:], username)

		go func() {
			newFollowings, err := s.GetNewFollowings(context.Background(), username)
			if err != nil {
				log.Printf("Error: track followings - %s\n", err.Error())
			}

			for _, followingUsername := range newFollowings {
				actions <- fmt.Sprintf("<b>%s</b> follows <b>%s</b>\n", username, followingUsername)
				err := s.db.AddAction(context.Background(), &models.Action{
					Time:           time.Now(),
					Type:           "follow",
					Username:       username,
					TargetUsername: followingUsername,
				})
				if err != nil {
					log.Printf("error adding action: %s", err.Error())
				}
			}
		}()
	}
}
