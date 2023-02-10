package models

type Tweet struct {
	ID             string `db:"id"`
	AuthorID       string `db:"author_id"`
	AuthorUsername string `db:"author_username"`
	Text           string `db:"text"`
}
