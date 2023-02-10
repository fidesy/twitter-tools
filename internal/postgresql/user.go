package postgresql

import (
	"context"
	"github.com/fidesy/twitter-tools/internal/models"
	"time"
)

func (p *PostgreSQL) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := p.db.GetContext(
		ctx,
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
	err := p.db.GetContext(
		ctx,
		&username,
		"SELECT username FROM users WHERE is_tracked ORDER BY latest_ping LIMIT 1",
	)
	if err != nil {
		return "", err
	}

	_, err = p.db.ExecContext(ctx, "UPDATE users SET latest_ping=$1 WHERE username=LOWER($2)", time.Now().UTC(), username)
	if err != nil {
		return "", err
	}

	return username, nil
}

func (p *PostgreSQL) AddUser(ctx context.Context, user *models.User) error {
	_, err := p.db.NamedExecContext(
		ctx,
		"INSERT INTO users VALUES(:id, LOWER(:username), :name, :description, :profile_image_url, :following, :followers, :tweets, :is_tracked, :latest_ping)",
		user,
	)
	return err
}

func (p *PostgreSQL) UpdateUser(ctx context.Context, user *models.User) error {
	_, err := p.db.NamedExecContext(
		ctx,
		"UPDATE users SET is_tracked=:is_tracked, latest_ping=:latest_ping WHERE username=LOWER(:username)",
		user,
	)
	return err
}

func (p *PostgreSQL) DeleteUser(ctx context.Context, username string) error {
	_, err := p.db.ExecContext(
		ctx,
		"DELETE FROM users WHERE username=LOWER($1);",
		username,
	)
	return err
}
