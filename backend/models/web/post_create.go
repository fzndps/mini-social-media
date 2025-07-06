package web

type PostCreateRequest struct {
	UserId   int    `json:"user_id" validate:"required,min=1"`
	Content  string `json:"content" validate:"required,min=1,max=1000"`
	ImageURL string `json:"image_url" validate:"required"`
}

type PostCreateResponse struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Content   string `json:"content"`
	ImageURL  string `json:"image_url"`
	CreatedAt string `json:"created_at"`
}
