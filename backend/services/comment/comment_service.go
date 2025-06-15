package comment

import (
	"context"

	"github.com/fzndps/mini-social-media/backend/models/domain"
	"github.com/fzndps/mini-social-media/backend/models/web"
)

type CommentService interface {
	CreateComment(ctx context.Context, request web.CommentCreateRequest) web.CommentResponse
	FindPostWithCommentsById(ctx context.Context, postId int) (domain.PostWithComments, error)
	UpdateComment(ctx context.Context, commentId int, userId int, request web.UpdateCommentRequest) web.CommentResponse
	Delete(ctx context.Context, commentId int, postId int, userId int)
}
