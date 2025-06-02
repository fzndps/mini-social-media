package domain

import "time"

type User struct {
	Id         int
	Username   string
	Email      string
	Password   string
	Created_at time.Time
}
