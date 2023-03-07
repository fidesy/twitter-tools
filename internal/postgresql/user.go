package postgresql

import (
	"context"
	"github.com/fidesy/twitter-tools/internal/models"
	"github.com/georgysavva/scany/v2/pgxscan"
	"time"
)

func (p *PostgreSQL) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User

	err := pgxscan.Get(
		ctx,
		p.pool,
		&user,
		"SELECT * FROM users WHERE username=LOWER($1)",
		username,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *PostgreSQL) GetUsernameToPing(ctx context.Context) (string, error) {
	var username string
	err := pgxscan.Get(
		ctx,
		p.pool,
		&username,
		"SELECT username FROM users WHERE is_tracked ORDER BY latest_ping LIMIT 1",
	)
	if err != nil {
		return "", err
	}

	_, err = p.pool.Exec(ctx, "UPDATE users SET latest_ping=$1 WHERE username=LOWER($2)", time.Now().UTC(), username)
	if err != nil {
		return "", err
	}

	return username, nil
}

func (p *PostgreSQL) AddUser(ctx context.Context, user *models.User) error {
	_, err := p.pool.Exec(
		ctx,
		"INSERT INTO users VALUES($1, LOWER($2), $3, $4, $5, $6, $7, $8, $9, $10)",
		user.ID, user.Username, user.Name, user.Description, user.ProfileImageURL,
		user.Following, user.Followers, user.Tweets, user.IsTracked, user.LatestPing,
	)
	return err
}

func (p *PostgreSQL) UpdateUser(ctx context.Context, user *models.User) error {
	_, err := p.pool.Exec(
		ctx,
		"UPDATE users SET is_tracked=$1, latest_ping=$2 WHERE username=LOWER($3)",
		user.IsTracked,
		user.LatestPing,
		user.Username,
	)
	return err
}

func (p *PostgreSQL) DeleteUser(ctx context.Context, username string) error {
	_, err := p.pool.Exec(
		ctx,
		"DELETE FROM users WHERE username=LOWER($1);",
		username,
	)
	return err
}
