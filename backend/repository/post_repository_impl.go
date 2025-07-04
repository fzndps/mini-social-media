package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/fzndps/mini-social-media/backend/helper"
	"github.com/fzndps/mini-social-media/backend/models/domain"
)

type PostRepositoryImpl struct {
	DB *sql.DB
}

func NewPostRepository(DB *sql.DB) PostRepository {
	return &PostRepositoryImpl{
		DB: DB,
	}
}

func (repository *PostRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post {
	SQL := "insert into posts (user_id, content, image_url) values(?, ?, ?)"

	result, err := tx.ExecContext(ctx, SQL, post.UserId, post.Content, post.ImageURL)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	post.Id = int(id)
	log.Println(post)
	return post
}

func (repository *PostRepositoryImpl) FindAll(ctx context.Context) []domain.Post {
	SQL := `SELECT 
						p.id, p.user_id, p.content, p.image_url, p.created_at, 
						u.id, u.username,
						(SELECT COUNT(*) FROM comments c WHERE c.post_id = p.id) AS comment_count
					FROM posts p
					JOIN users u ON p.user_id = u.id
					ORDER BY p.created_at DESC`

	rows, err := repository.DB.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	log.Println("Ini isi dari rows repository", rows)
	defer rows.Close()

	var posts []domain.Post

	for rows.Next() {
		post := domain.Post{}
		post.User = domain.UserPost{}
		err := rows.Scan(&post.Id, &post.UserId, &post.Content, &post.ImageURL, &post.CreatedAt, &post.User.Id, &post.User.Username, &post.CommentCount)
		if err != nil {
			log.Println("Scan error:", err)
			helper.PanicIfError(err)
		}

		log.Println("post ditemukan:", post)
		posts = append(posts, post)
	}

	return posts
}

func (repository *PostRepositoryImpl) DeletePost(ctx context.Context, tx *sql.Tx, postId int, userId int) {
	SQL := "delete from posts where id = ? and user_id = ?"
	_, err := tx.ExecContext(ctx, SQL, postId, userId)
	helper.PanicIfError(err)
}

func (repository *PostRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, postId int) (domain.Post, error) {
	SQL := "select id, user_id, content, created_at from posts where id = ?"
	rows := tx.QueryRowContext(ctx, SQL, postId)

	post := domain.Post{}
	err := rows.Scan(&post.Id, &post.UserId, &post.Content, &post.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return post, errors.New("Postingan tidak ditemukan")
		}
		return post, nil
	}

	return post, nil
}
