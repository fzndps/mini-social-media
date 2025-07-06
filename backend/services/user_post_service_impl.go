package services

import (
	"context"
	"database/sql"
	"log"

	"github.com/fzndps/mini-social-media/backend/helper"
	"github.com/fzndps/mini-social-media/backend/models/web"
	"github.com/fzndps/mini-social-media/backend/repository"
)

type UserPostServiceImpl struct {
	UserPostRepository repository.UserPostRepository
	DB                 *sql.DB
}

func NewUserPostService(UserPostService repository.UserPostRepository, DB *sql.DB) UserPostService {
	return &UserPostServiceImpl{
		UserPostRepository: UserPostService,
		DB:                 DB,
	}
}

func (service *UserPostServiceImpl) UserPostProfile(ctx context.Context, userId int) web.UserPostResponse {
	userPost := service.UserPostRepository.UserPostProfile(ctx, userId)
	log.Println(userPost)
	return helper.ToUserPostResponses(userPost)
}
