package service

import (
	"fmt"
	"github.com/fidesy/twitter-tools/internal/models"
	"github.com/g8rswimmer/go-twitter/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	db     *gorm.DB
}

func New(config *Config) (*Service, error) {
	s := &Service{}

	proxy, err := url.Parse(config.RawProxy)
	if err != nil {
		return nil, err
	}

	cli := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}

	client := &twitter.Client{
		Authorizer: authorize{Token: config.BearerToken},
		Client:     cli,
		Host:       "https://api.twitter.com",
	}
	s.client = client

	err = s.configureDatabase(config.DBURL)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Service) configureDatabase(dbURL string) error {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(models.Tweet{})
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Following{})

	s.db = db

	return nil
}
