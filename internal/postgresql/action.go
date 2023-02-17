package postgresql

import (
	"context"
	"github.com/fidesy/twitter-tools/internal/models"
	"time"
)

func (p *PostgreSQL) AddAction(ctx context.Context, action *models.Action) error {
	_, err := p.db.NamedExecContext(
		ctx,
		"INSERT INTO actions(time, type, username, target_username) VALUES(:time, :type, LOWER(:username), LOWER(:target_username))",
		action,
	)

	return err
}

func (p *PostgreSQL) GetTopFollowings(ctx context.Context, duration time.Duration, limit int) ([]models.TopFollowing, error) {
	var topFollowing []models.TopFollowing
	err := p.db.SelectContext(
		ctx,
		&topFollowing,
		`SELECT target_username, COUNT(*) AS amount FROM actions WHERE type='follow' AND time > $1 GROUP BY target_username HAVING COUNT(*) >= 2 ORDER BY COUNT(*) DESC LIMIT $2`,
		time.Now().Add(-duration).UTC(),
		limit,
	)
	if err != nil {
		return nil, err
	}

	return topFollowing, nil
}

func (p *PostgreSQL) GetTopFollowers(ctx context.Context, username string, duration time.Duration, limit int) ([]string, error) {
	var topFollowers []string
	err := p.db.SelectContext(
		ctx,
		&topFollowers,
		`SELECT actions.username FROM actions JOIN users ON (actions.username=users.username) WHERE actions.type='follow' AND actions.target_username=LOWER($1) AND actions.time > $2 ORDER BY users.followers DESC LIMIT $3`,
		username,
		time.Now().Add(-duration).UTC(),
		limit,
	)
	if err != nil {
		return nil, err
	}

	return topFollowers, nil
}
