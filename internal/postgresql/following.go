package postgresql

import (
	"context"
	"github.com/fidesy/twitter-tools/internal/models"
)

func (p *PostgreSQL) GetFollowingsByUsername(ctx context.Context, username string) ([]models.Following, error) {
	var followings []models.Following
	err := p.db.SelectContext(
		ctx,
		&followings,
		"SELECT * FROM followings WHERE username=LOWER($1)",
		username,
	)
	if err != nil {
		return nil, err
	}

	return followings, nil
}

func (p *PostgreSQL) AddFollowing(ctx context.Context, following *models.Following) error {
	_, err := p.db.NamedExecContext(
		ctx,
		"INSERT INTO followings(user_id, username, following_id, following_username) VALUES(:user_id, LOWER(:username), :following_id, LOWER(:following_username))",
		following,
	)
	return err
}
