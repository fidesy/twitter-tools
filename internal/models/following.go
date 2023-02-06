package models

type Following struct {
	ID                int    `db:"id"`
	UserID            string `db:"user_id"`
	Username          string `db:"username"`
	FollowingID       string `db:"following_id"`
	FollowingUsername string `db:"following_username"`
}
