package postgresql

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOpenPostgres(t *testing.T) {
	db := New()
	err := db.Open(context.Background(), os.Getenv("DB_URL"))
	assert.Nil(t, err)
	db.Close()
}
