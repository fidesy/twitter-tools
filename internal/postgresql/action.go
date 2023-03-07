package postgresql

import (
	"context"
	"github.com/fidesy/twitter-tools/internal/models"
	"github.com/georgysavva/scany/v2/pgxscan"
	"time"
)

func (p *PostgreSQL) AddAction(ctx context.Context, action *models.Action) error {
	_, err := p.pool.Exec(
		ctx,
		"INSERT INTO actions(time, type, username, target_username) VALUES($1, $2, LOWER($3), LOWER($4))",
		action.Time, action.Type,
		action.Username, action.TargetUsername,
	)

	return err
}

func (p *PostgreSQL) GetTopFollowings(ctx context.Context, duration time.Duration, limit int) ([]models.TopFollowing, error) {
	var topFollowing []models.TopFollowing
	err := pgxscan.Select(
		ctx,
		p.pool,
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
	err := pgxscan.Select(
		ctx,
		p.pool,
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
