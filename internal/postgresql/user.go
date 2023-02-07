package postgresql

import (
	"context"
	"github.com/fidesy/twitter-tools/internal/models"
)

const (
	selectUserByUsernameQuery   = "SELECT * FROM users WHERE username=LOWER($1)"
	selectUsersForTrackingQuery = "SELECT username FROM users WHERE is_tracked is true"
	insertUserQuery             = "INSERT INTO users VALUES(:id, LOWER(:username), :name, :description, :profile_image_url, :following, :followers, :tweets, :is_tracked)"
	updateUserQuery             = "UPDATE users SET is_tracked=:is_tracked WHERE username=LOWER(:username)"
	deleteUserQuery             = "DELETE FROM users WHERE username=LOWER($1);"
)

func (p *PostgreSQL) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user = new(models.User)
	err := p.db.GetContext(ctx, &user, selectUserByUsernameQuery, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (p *PostgreSQL) GetUsersForTracking(ctx context.Context) ([]string, error) {
	var usernames []string
	err := p.db.SelectContext(ctx, &usernames, selectUsersForTrackingQuery)
	if err != nil {
		return nil, err
	}

	return usernames, nil
}

func (p *PostgreSQL) AddUser(ctx context.Context, user *models.User) error {
	_, err := p.db.NamedExecContext(ctx, insertUserQuery, user)
	return err
}

func (p *PostgreSQL) UpdateUserTrackField(ctx context.Context, user *models.User) error {
	_, err := p.db.NamedExecContext(ctx, updateUserQuery, user)
	return err
}

func (p *PostgreSQL) DeleteUser(ctx context.Context, username string) error {
	_, err := p.db.ExecContext(ctx, deleteUserQuery, username)
	return err
}
