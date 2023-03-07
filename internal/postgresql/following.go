package postgresql

import (
	"context"
	"github.com/fidesy/twitter-tools/internal/models"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (p *PostgreSQL) GetFollowingsByUsername(ctx context.Context, username string) ([]models.Following, error) {
	var followings []models.Following
	err := pgxscan.Select(
		ctx,
		p.pool,
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
	_, err := p.pool.Exec(
		ctx,
		"INSERT INTO followings(user_id, username, following_id, following_username) VALUES($1, LOWER($2), $3, LOWER($4))",
		following.UserID, following.Username,
		following.FollowingID, following.FollowingUsername,
	)
	return err
}
