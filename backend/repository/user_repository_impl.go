package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/fzndps/mini-social-media/backend/helper"
	"github.com/fzndps/mini-social-media/backend/models/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "insert into users (username, email, password) value(?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, user.Username, user.Email, user.Password)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)
	return user
}

func (repository *UserRepositoryImpl) LoginByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error) {
	SQL := "SELECT id, username, email, password FROM users WHERE username = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, SQL, username)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if err != nil {
		return user, err
	}

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		return user, err
	} else {
		return user, errors.New("user not found")
	}
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error) {
	SQL := "select username, email from users where username = ?"
	rows, err := tx.QueryContext(ctx, SQL, username)
	helper.PanicIfError(err)

	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Username, &user.Email)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (repository *UserRepositoryImpl) IsUsernameExists(ctx context.Context, tx *sql.Tx, username string) bool {
	SQL := "select count(*) from users where username = ?"
	var count int
	err := tx.QueryRowContext(ctx, SQL, username).Scan(&count)
	helper.PanicIfError(err)
	return count > 0
}

func (repository *UserRepositoryImpl) IsEmailExists(ctx context.Context, tx *sql.Tx, email string) bool {
	SQL := "select count(*) from users where email = ?"
	var count int
	err := tx.QueryRowContext(ctx, SQL, email).Scan(&count)
	helper.PanicIfError(err)
	return count > 0
}
