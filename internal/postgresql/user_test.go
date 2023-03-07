package postgresql

import (
	"context"
	"github.com/fidesy/twitter-tools/internal/models"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	userDB = New()
	user   = &models.User{Username: "danila"}
)

func TestUser_OpenDatabase(t *testing.T) {
	err := userDB.Open(context.Background(), os.Getenv("DB_URL"))
	assert.Nil(t, err)
}

func TestUser_Add(t *testing.T) {
	err := userDB.AddUser(context.Background(), user)
	assert.Nil(t, err)
}

func TestUser_GetUserByUsername(t *testing.T) {
	u, err := userDB.GetUserByUsername(context.Background(), user.Username)
	assert.Nil(t, err)

	t.Log("User:", u)
}

func TestUser_UpdateUserTrackField(t *testing.T) {
	user.IsTracked = true
	err := userDB.UpdateUser(context.Background(), user)
	assert.Nil(t, err)

	u, err := userDB.GetUserByUsername(context.Background(), user.Username)
	assert.Nil(t, err)

	assert.Equal(t, user.IsTracked, u.IsTracked)
}

func TestUser_Delete(t *testing.T) {
	err := userDB.DeleteUser(context.Background(), user.Username)
	assert.Nil(t, err)

	userDB.Close()
}
