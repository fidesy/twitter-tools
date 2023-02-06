package models

type TopFollowings struct {
	Username string `db:"target_username"`
	Amount   int    `db:"amount"`
}
