package services

import (
	"context"
	"database/sql"
	"log"

	"github.com/fzndps/mini-social-media/backend/helper"
	"github.com/fzndps/mini-social-media/backend/models/domain"
	"github.com/fzndps/mini-social-media/backend/models/web"
	"github.com/fzndps/mini-social-media/backend/repository"
	"github.com/go-playground/validator/v10"
)

type PostServiceImpl struct {
	PostRepository repository.PostRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewPostService(postRepository repository.PostRepository, DB *sql.DB, validate *validator.Validate) PostService {
	return &PostServiceImpl{
		PostRepository: postRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *PostServiceImpl) Create(ctx context.Context, request web.PostCreateRequest) web.PostCreateResponse {
	log.Println("Service Create start =>", request)
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	log.Println("Service Create start =>", tx)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	post := domain.Post{
		UserId:   request.UserId,
		Content:  request.Content,
		ImageURL: request.ImageURL,
	}

	log.Println("Calling repository.Create with post =>", post)
	post = service.PostRepository.Create(ctx, tx, post)
	log.Println("Post created successfully =>", post)
	return helper.ToCreatePostResponse(post)
}

func (service *PostServiceImpl) FindAll(ctx context.Context) []web.FindAllPostResponses {
	posts := service.PostRepository.FindAll(ctx)
	log.Println("✅Isi ini posts dari service", posts)
	return helper.ToPostResponses(posts)
}
