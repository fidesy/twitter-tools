package postgresql

import (
	"context"
	"github.com/fidesy/twitter-tools/internal/models"
	"strings"
)

const (
	selectUserByUsernameQuery = "SELECT * FROM users WHERE username=$1"
	insertUserQuery           = "INSERT INTO users VALUES(:id, :username, :name, :description, :profile_image_url, :following, :followers, :tweets, :is_tracked)"
	updateUserQuery           = "UPDATE users SET is_tracked=:is_tracked WHERE username=:username"
)

func (p *PostgreSQL) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user = new(models.User)
	err := p.db.GetContext(ctx, &user, selectUserByUsernameQuery, strings.ToLower(username))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (p *PostgreSQL) AddUser(ctx context.Context, user *models.User) error {
	user.Username = strings.ToLower(user.Username)
	_, err := p.db.NamedExecContext(ctx, insertUserQuery, user)
	return err
}

func (p *PostgreSQL) UpdateUserTrackField(ctx context.Context, user *models.User) error {
	_, err := p.db.NamedExecContext(ctx, updateUserQuery, user)
	return err
}
