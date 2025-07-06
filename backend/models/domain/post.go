package domain

type Post struct {
	Id           int
	UserId       int
	Content      string
	ImageURL     string
	CreatedAt    string
	CommentCount int
	User         UserPost
}

type UserPost struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}
