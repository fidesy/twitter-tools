package models

type TopFollowing struct {
	Username string `db:"target_username"`
	Amount   int    `db:"amount"`
}
