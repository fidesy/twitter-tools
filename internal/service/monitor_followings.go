package service

import (
	"context"
	"fmt"
	"log"
	"time"
)

func (s *Service) MonitorFollowings(actions chan<- string, usernames []string) {
	for _ = range time.Tick(time.Minute) {
		username := usernames[0]
		usernames = append(usernames[1:], username)

		go func() {
			newFollowings, delFollowings, err := s.GetFollowingActions(context.Background(), username)
			if err != nil {
				log.Printf("Error: monitor followings - %s\n", err.Error())
			}

			for _, followingUsername := range newFollowings {
				actions <- fmt.Sprintf("<b>%s</b> follows <b>%s</b>\n", username, followingUsername)
			}

			for _, followingUsername := range delFollowings {
				actions <- fmt.Sprintf("<b>%s</b> unfollows <b>%s</b>", username, followingUsername)
			}
		}()
	}
}
