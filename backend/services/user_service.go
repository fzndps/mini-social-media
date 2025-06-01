package services

import (
	"context"

	"github.com/fzndps/mini-social-media/backend/models/web"
)

type UserService interface {
	Register(ctx context.Context, request web.UserRegisterRequest) web.UserRegisterResponse
	Login(ctx context.Context, request web.UserLoginRequest) (web.UserLoginResponse, error)
	FindByUsername(ctx context.Context, username string) web.UserRegisterResponse
}
