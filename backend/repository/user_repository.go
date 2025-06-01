package repository

import (
	"context"
	"database/sql"

	"github.com/fzndps/mini-social-media/backend/models/domain"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	LoginByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error)
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error)
	IsUsernameExists(ctx context.Context, tx *sql.Tx, username string) bool
	IsEmailExists(ctx context.Context, tx *sql.Tx, email string) bool
}
