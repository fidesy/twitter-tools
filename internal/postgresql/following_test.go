package postgresql

import (
	"context"
	"github.com/fidesy/twitter-tools/internal/models"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	followingDB = New()
)

func TestOpenFollowingDB(t *testing.T) {
	err := followingDB.Open(context.Background(), os.Getenv("DB_URL"))
	assert.Nil(t, err)
}

func TestGetFollowingsByUsername(t *testing.T) {
	followings, err := followingDB.GetFollowingsByUsername(context.Background(), "danila")

	assert.Nil(t, err)

	t.Log("Followings:", followings)
}

func TestAddFollowing(t *testing.T) {
	err := followingDB.AddFollowing(context.Background(), &models.Following{Username: "max", FollowingUsername: "dan"})
	assert.Nil(t, err)

	followingDB.Close()
}