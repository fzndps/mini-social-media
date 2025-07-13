package domain

import (
	"database/sql"
	"time"
)

type User struct {
	Id         int
	Username   string
	Email      string
	Password   string
	Created_at time.Time
}

type UserwithPost struct {
	Id       int
	Username string
	Posts    []UserProfilePost
}

type UserProfilePost struct {
	PostId       sql.NullInt32
	Content      sql.NullString
	ImageURL     sql.NullString
	CreatedAt    sql.NullString
	CommentCount sql.NullInt32
}
