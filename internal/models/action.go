package models

import "time"

type Action struct {
	ID             int       `db:"id"`
	Time           time.Time `db:"time"`
	Type           string    `db:"type"`
	Username       string    `db:"username"`
	TargetUsername string    `db:"target_username"`
}
