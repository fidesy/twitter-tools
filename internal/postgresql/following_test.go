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

func TestFollowing_OpenDatabase(t *testing.T) {
	err := followingDB.Open(context.Background(), os.Getenv("DB_URL"))
	assert.Nil(t, err)
}

func TestFollowing_GetFollowingsByUsername(t *testing.T) {
	followings, err := followingDB.GetFollowingsByUsername(context.Background(), "danila")

	assert.Nil(t, err)

	t.Log("Followings:", followings)
}

func TestFollowing_Add(t *testing.T) {
	err := followingDB.AddFollowing(context.Background(), &models.Following{
		UserID:            "1",
		Username:          "max",
		FollowingID:       "2",
		FollowingUsername: "dan"})
	assert.Nil(t, err)

	followingDB.Close()
}
