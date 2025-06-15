package comment

import (
	"context"
	"database/sql"
	"log"

	"github.com/fzndps/mini-social-media/backend/exception"
	"github.com/fzndps/mini-social-media/backend/helper"
	"github.com/fzndps/mini-social-media/backend/models/domain"
	"github.com/fzndps/mini-social-media/backend/models/web"
	"github.com/fzndps/mini-social-media/backend/repository"
	"github.com/fzndps/mini-social-media/backend/repository/comment"
	"github.com/go-playground/validator/v10"
)

type CommentServiceImpl struct {
	CommentRepository comment.CommentRepository
	PostRepository    repository.PostRepository
	UserRepository    repository.UserRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewCommentService(commentRepository comment.CommentRepository, postRepository repository.PostRepository, userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) CommentService {
	return &CommentServiceImpl{
		CommentRepository: commentRepository,
		PostRepository:    postRepository,
		UserRepository:    userRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *CommentServiceImpl) CreateComment(ctx context.Context, request web.CommentCreateRequest) web.CommentResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	comment := domain.CommentWitUser{
		Content: request.Content,
		UserId:  request.UserId,
		PostId:  request.PostId,
	}

	comment = service.CommentRepository.Create(ctx, tx, comment)
	user, err := service.UserRepository.FindById(ctx, tx, comment.UserId)
	helper.PanicIfError(err)

	comment.User = domain.UserCommentInfo{
		Id:       user.Id,
		Username: user.Username,
	}

	return helper.ToCommentResponse(comment)
}

func (service *CommentServiceImpl) FindPostWithCommentsById(ctx context.Context, postId int) (domain.PostWithComments, error) {
	post, err := service.CommentRepository.FindPostWithCommentsById(ctx, postId)
	if err != nil {
		panic(exception.NewNotFoundError("Post not found"))
	}

	return post, nil
}

func (service *CommentServiceImpl) UpdateComment(ctx context.Context, commentId int, userId int, request web.UpdateCommentRequest) web.CommentResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	log.Println("Ini tx dari service :", tx)

	existingComment, err := service.CommentRepository.FindById(ctx, tx, commentId)
	if err != nil {
		panic(exception.NewNotFoundError("Comment not found"))
	}

	if existingComment.UserId != userId {
		panic(exception.NewUnauthorizedError("You do not have the ability to edit comments made by other people."))
	}

	existingComment.Content = request.Content

	updateComment := service.CommentRepository.Update(ctx, tx, existingComment)

	user, err := service.UserRepository.FindById(ctx, tx, updateComment.UserId)
	if err != nil {
		panic(exception.NewNotFoundError("user id not found"))
	}

	updateComment.User = domain.UserCommentInfo{
		Id:       user.Id,
		Username: user.Username,
	}

	return helper.ToCommentResponse(updateComment)
}

func (service *CommentServiceImpl) Delete(ctx context.Context, commentId int, postId int, userId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	existingComment, err := service.CommentRepository.FindById(ctx, tx, commentId)
	if err != nil {
		panic(exception.NewNotFoundError("Comment not found"))
	}

	if existingComment.PostId != postId {
		panic(exception.NewNotFoundError("Comment not found in this post"))
	}

	if existingComment.UserId != userId {
		panic(exception.NewUnauthorizedError("You do not have permission to delete this comment"))
	}

	service.CommentRepository.Delete(ctx, tx, commentId)
}
