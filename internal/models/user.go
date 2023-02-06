package models

type User struct {
	ID              string `db:"id"`
	Username        string `db:"username"`
	Name            string `db:"name"`
	Description     string `db:"description"`
	ProfileImageURL string `db:"profile_image_url"`
	Following       int    `db:"following"`
	Followers       int    `db:"followers"`
	Tweets          int    `db:"tweets"`
	IsTracked       bool   `db:"is_tracked"`
}
