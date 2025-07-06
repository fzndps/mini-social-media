package web

type CommentResponse struct {
	Id        int      `json:"id"`
	Content   string   `json:"content"`
	PostId    int      `json:"post_id"`
	CreatedAt string   `json:"created_at"`
	User      UserInfo `json:"user_info"`
}

type UserInfo struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type UpdateCommentRequest struct {
	// Id      int    `json:"id"`
	Content string `json:"content" validate:"required,min=1"`
	// UserId  int    `json:"user_id"`
	// PostId  int    `json:"post_id"`
}

type CommentCreateRequest struct {
	Content string `json:"content"`
	UserId  int    `json:"user_id"`
	PostId  int    `json:"post_id"`
}
