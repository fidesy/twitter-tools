package service

import (
	"context"
	"github.com/g8rswimmer/go-twitter/v2"
	"log"
)

func (s *Service) GetLatestTweets(ctx context.Context, username string) ([]*twitter.TweetObj, error) {
	userID, err := s.GetIDByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.UserTweetTimeline(ctx, userID, twitter.UserTweetTimelineOpts{})
	if err != nil {
		return nil, err
	}

	log.Printf("GetLatestTweets rate limit %d\n", resp.RateLimit.Remaining)

	return resp.Raw.Tweets, nil
}
