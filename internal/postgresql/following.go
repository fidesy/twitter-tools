package postgresql

import (
	"context"
	"github.com/fidesy/twitter-tools/internal/models"
	"strings"
)

const (
	selectFollowingsQuery = "SELECT * FROM followings WHERE username=$1"
	insertFollowingQuery  = "INSERT INTO followings(user_id, username, following_id, following_username) VALUES(:user_id, :username, :following_id, :following_username)"
)

func (p *PostgreSQL) GetFollowingsByUsername(ctx context.Context, username string) ([]models.Following, error) {
	followings := []models.Following{}

	err := p.db.SelectContext(ctx, &followings, selectFollowingsQuery, strings.ToLower(username))
	if err != nil {
		return nil, err
	}

	return followings, nil
}

func (p *PostgreSQL) AddFollowing(ctx context.Context, following *models.Following) error {
	_, err := p.db.NamedExecContext(ctx, insertFollowingQuery, following)
	return err
}
