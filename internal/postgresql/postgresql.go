package postgresql

import (
	"context"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgreSQL struct {
	pool *pgxpool.Pool
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
    id SERIAL PRIMARY KEY,
    time TIMESTAMP,
    type TEXT,
    username TEXT,
    target_username TEXT
);`

func (p *PostgreSQL) Open(ctx context.Context, dbURL string) error {
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return err
	}

	if err = pool.Ping(ctx); err != nil {
		return err
	}

	if _, err = pool.Exec(ctx, initSchemas); err != nil {
		return err
	}

	p.pool = pool
	return nil
}

func (p *PostgreSQL) Close() {
	p.pool.Close()
}
