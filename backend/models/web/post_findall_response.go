package web

type FindAllPostResponses struct {
	Id        int          `json:"id"`
	Content   string       `json:"content"`
	ImageURL  string       `json:"image_url"`
	CreatedAt string       `json:"created_at"`
	User      UserResponse `json:"user"`
}

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}
