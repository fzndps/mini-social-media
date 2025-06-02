package helper

import (
	"github.com/fzndps/mini-social-media/backend/models/domain"
	"github.com/fzndps/mini-social-media/backend/models/web"
)

func ToUserResponse(user domain.User) web.UserRegisterResponse {
	return web.UserRegisterResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}
}

func ToUserLoginResponse(user domain.User, token string) web.UserLoginResponse {
	return web.UserLoginResponse{
		TokenType:   "Bearer ",
		AccessToken: token,
		ExpiresIn:   3600,
		User: web.UserRegisterResponse{
			Id:       user.Id,
			Username: user.Username,
			Email:    user.Email,
		},
	}
}
