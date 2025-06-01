package web

type UserRegisterRequest struct {
	Username string `validate:"required,max=50,min=1" json:"username"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,max=255,min=8" json:"password"`
}
