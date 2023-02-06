package service

import (
	"context"
	"github.com/g8rswimmer/go-twitter/v2"
	"log"
)

func (s *Service) GetLatestTweets(ctx context.Context, username string) ([]*twitter.TweetObj, error) {
	user, err := s.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.UserTweetTimeline(ctx, user.ID, twitter.UserTweetTimelineOpts{})
	if err != nil {
		return nil, err
	}

	log.Printf("GetLatestTweets rate limit %d\n", resp.RateLimit.Remaining)

	return resp.Raw.Tweets, nil
}
