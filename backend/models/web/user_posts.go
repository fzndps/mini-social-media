package web

import "database/sql"

type UserPostResponse struct {
	Id       int
	Username string
	Posts    []UserPostProfile
}

type UserPostProfile struct {
	PostId    sql.NullInt32
	Content   sql.NullString
	ImageURL  sql.NullString
	CreatedAt sql.NullString
}
