package comment

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/fzndps/mini-social-media/backend/helper"
	"github.com/fzndps/mini-social-media/backend/models/domain"
)

type CommentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(DB *sql.DB) CommentRepository {
	return &CommentRepositoryImpl{
		DB: DB,
	}
}

func (repository *CommentRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, comment domain.CommentWitUser) domain.CommentWitUser {
	SQL := "insert into comments(content, user_id, post_id) values(?, ?, ?)"

	result, err := tx.ExecContext(ctx, SQL, comment.Content, comment.UserId, comment.PostId)
	log.Println("Ini result exec: ", result)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	comment.Id = int(id)

	log.Println("Ini comment repo : ", comment)

	row := tx.QueryRowContext(ctx, "SELECT created_at FROM comments WHERE id = ?", comment.Id)
	err = row.Scan(&comment.CreatedAt)
	helper.PanicIfError(err)

	return comment
}

func (repository *CommentRepositoryImpl) FindPostWithCommentsById(ctx context.Context, postId int) (domain.PostWithComments, error) {
	tx, err := repository.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQLPost := `SELECT p.id, p.user_id, u.username, p.content, p.image_url, p.created_at
		FROM posts p
		JOIN users u ON p.user_id = u.id
		WHERE p.id = ?`

	var post domain.PostWithComments
	err = tx.QueryRowContext(ctx, SQLPost, postId).Scan(&post.Id, &post.UserId, &post.Username, &post.Content, &post.ImageURL, &post.CreatedAt)
	if err != nil {
		return post, err
	}

	SQLComments := `SELECT c.id, c.user_id, u.username, c.content, c.created_at
		FROM comments c
		JOIN users u ON c.user_id = u.id
		WHERE c.post_id = ?
		ORDER BY c.created_at ASC`

	rows, err := tx.QueryContext(ctx, SQLComments, postId)
	helper.PanicIfError(err)
	defer rows.Close()

	var comments []domain.CommentWitUser
	for rows.Next() {
		comment := domain.CommentWitUser{}
		err := rows.Scan(&comment.Id, &comment.UserId, &comment.User.Username, &comment.Content, &comment.CreatedAt)
		helper.PanicIfError(err)
		comments = append(comments, comment)
	}

	post.Comments = comments

	return post, nil
}

func (repository *CommentRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, comment domain.CommentWitUser) domain.CommentWitUser {
	SQL := "update comments set content = ? where id = ? and user_id = ?"

	res, err := tx.ExecContext(ctx, SQL, comment.Content, comment.Id, comment.UserId)
	helper.PanicIfError(err)

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		panic(errors.New("no comment updated"))
	}

	return comment
}

func (repository *CommentRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, commentId int) {
	SQL := "delete from comments where id = ?"
	_, err := tx.ExecContext(ctx, SQL, commentId)
	helper.PanicIfError(err)
}

func (repository *CommentRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, commentId int) (domain.CommentWitUser, error) {
	query := "SELECT id, content, user_id, post_id, created_at FROM comments WHERE id = ?"
	row := tx.QueryRowContext(ctx, query, commentId)

	comment := domain.CommentWitUser{}
	err := row.Scan(&comment.Id, &comment.Content, &comment.UserId, &comment.PostId, &comment.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return comment, errors.New("comment tidak ditemukan")
		}
		return comment, err
	}

	return comment, nil
}
