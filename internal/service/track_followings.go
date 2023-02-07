package service

import (
	"context"
	"fmt"
	"github.com/fidesy/twitter-tools/internal/models"
	"log"
	"time"
)

func (s *Service) TrackFollowings(actions chan<- string) {
	var (
		usernames = []string{}
		err       error
	)

	for _ = range time.Tick(time.Minute) {
		if len(usernames) == 0 {
			usernames, err = s.db.GetUsersForTracking(context.Background())
			if err != nil {
				log.Println("error fetching new users for tracking:", err.Error())
			}
			log.Println("Users to track:", len(usernames))
		}

		username := usernames[0]
		usernames = usernames[1:]

		go func() {
			newFollowings, err := s.GetNewFollowings(context.Background(), username)
			if err != nil {
				log.Printf("Error: track followings - %s\n", err.Error())
			}

			for _, followingUsername := range newFollowings {
				actions <- fmt.Sprintf(
					`<a href="https://twitter.com/%s"><b>%s</b></a> follows <a href="https://twitter.com/%s"><b>%s</b></a>\n`,
					username, username, followingUsername, followingUsername)
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
