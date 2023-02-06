package postgresql

import (
	"context"
	"github.com/fidesy/twitter-tools/internal/models"
	"time"
)

const (
	insertActionQuery   = "INSERT INTO actions VALUES(:time, :type, :username, :target_username)"
	selectTopFollowings = `
		SELECT target_username, COUNT(*) AS amount FROM actions
		WHERE type='follow' AND time > $1
		GROUP BY target_username
		ORDER BY COUNT(*)
		`
	selectTopFollowers = `
    	SELECT DISTINCT actions.username FROM actions
    	JOIN users ON actions.username=users.username
		WHERE actions.type='follow' AND actions.target_username=$1 AND actions.time > $2
		ORDER BY users.followers DESC
		LIMIT 5
	`
)

func (p *PostgreSQL) AddAction(ctx context.Context, action *models.Action) error {
	_, err := p.db.NamedExecContext(ctx, insertActionQuery, action)
	return err
}

func (p *PostgreSQL) GetTopFollowings(ctx context.Context, duration time.Duration) ([]models.TopFollowings, error) {
	var topFollowing []models.TopFollowings
	err := p.db.SelectContext(ctx, &topFollowing, selectTopFollowings, time.Now().Add(-duration))
	if err != nil {
		return nil, err
	}

	return topFollowing, nil
}

func (p *PostgreSQL) GetTopFollowers(ctx context.Context, username string, duration time.Duration) ([]string, error) {
	var topFollowers []string
	err := p.db.SelectContext(ctx, &topFollowers, selectTopFollowers, username, time.Now().Add(-duration))
	if err != nil {
		return nil, err
	}

	return topFollowers, nil
}
