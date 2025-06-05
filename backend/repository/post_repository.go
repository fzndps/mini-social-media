package repository

import (
	"context"
	"database/sql"

	"github.com/fzndps/mini-social-media/backend/models/domain"
)

type PostRepository interface {
	Create(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post
	FindAll(ctx context.Context) []domain.Post
}
