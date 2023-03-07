package service

import (
	"context"
	"fmt"
	"github.com/fidesy/twitter-tools/internal/postgresql"
	"github.com/g8rswimmer/go-twitter/v2"
	"net/http"
	"net/url"
)

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

type Config struct {
	BearerToken string
	RawProxy    string
	DBURL       string
}

type Service struct {
	client *twitter.Client
	db     *postgresql.PostgreSQL
}

func New(config *Config) (*Service, error) {
	s := &Service{}

	cli := &http.Client{}
	if config.RawProxy != "" {
		proxy, err := url.Parse(config.RawProxy)
		if err != nil {
			return nil, err
		}

		cli = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxy),
			},
		}
	}

	client := &twitter.Client{
		Authorizer: authorize{Token: config.BearerToken},
		Client:     cli,
		Host:       "https://api.twitter.com",
	}
	s.client = client

	db := postgresql.New()
	if err := db.Open(context.Background(), config.DBURL); err != nil {
		return nil, err
	}
	s.db = db

	return s, nil
}

func (s *Service) CloseDatabase() {
	s.db.Close()
}
