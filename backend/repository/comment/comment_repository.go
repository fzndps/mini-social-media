package comment

import (
	"context"
	"database/sql"

	"github.com/fzndps/mini-social-media/backend/models/domain"
)

type CommentRepository interface {
	Create(ctx context.Context, tx *sql.Tx, comment domain.CommentWitUser) domain.CommentWitUser
	FindPostWithCommentsById(ctx context.Context, postId int) (domain.PostWithComments, error)
	Update(ctx context.Context, tx *sql.Tx, comment domain.CommentWitUser) domain.CommentWitUser
	Delete(ctx context.Context, tx *sql.Tx, commentId int)
	FindById(ctx context.Context, tx *sql.Tx, commentId int) (domain.CommentWitUser, error)
}
