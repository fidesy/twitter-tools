package postgresql

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	db *sqlx.DB
}

func New() *PostgreSQL {
	return &PostgreSQL{}
}

const initSchemas = `

CREATE TABLE IF NOT EXISTS users(
	id TEXT PRIMARY KEY,
	username TEXT UNIQUE,
	name TEXT,
	description TEXT,
	profile_image_url TEXT,
	following INT,
	followers INT,
	tweets INT,
	is_tracked BOOLEAN,
	latest_ping TIMESTAMP         
);

CREATE TABLE IF NOT EXISTS followings(
    id SERIAL PRIMARY KEY,
    user_id TEXT,
    username TEXT,
    following_id TEXT,
    following_username TEXT 
);

CREATE TABLE IF NOT EXISTS actions(
    time TIMESTAMP,
    type TEXT,
    username TEXT,
    target_username TEXT
);`

func (p *PostgreSQL) Open(ctx context.Context, dbURL string) error {
	db, err := sqlx.ConnectContext(ctx, "postgres", dbURL)
	if err != nil {
		return err
	}

	if err := db.PingContext(ctx); err != nil {
		return err
	}

	if _, err = db.ExecContext(ctx, initSchemas); err != nil {
		return err
	}

	p.db = db

	return nil
}

func (p *PostgreSQL) Close() {
	p.db.Close()
}
