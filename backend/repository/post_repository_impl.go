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
            posts.id, posts.user_id, posts.content, posts.image_url, posts.created_at, 
            users.id, users.username
        FROM posts
        JOIN users ON posts.user_id = users.id
        ORDER BY posts.created_at DESC`

	rows, err := repository.DB.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	log.Println("Ini isi dari rows repository", rows)
	defer rows.Close()

	var posts []domain.Post

	for rows.Next() {
		post := domain.Post{}
		post.User = domain.UserPost{}
		err := rows.Scan(&post.Id, &post.UserId, &post.Content, &post.ImageURL, &post.CreatedAt, &post.User.Id, &post.User.Username)
		if err != nil {
			log.Println("Scan error:", err)
			helper.PanicIfError(err)
		}

		log.Println("post ditemukan:", post)
		posts = append(posts, post)
	}

	return posts
}

func (repository *PostRepositoryImpl) FindById(ctx context.Context, postId int) (domain.Post, error) {
	SQL := "select id, user_id, content, created_at from posts where id = ?"
	rows, err := repository.DB.QueryContext(ctx, SQL, postId)
	helper.PanicIfError(err)

	defer rows.Close()

	post := domain.Post{}

	if rows.Next() {
		err := rows.Scan(&post.Id, &post.UserId, &post.Content, &post.CreatedAt)
		helper.PanicIfError(err)
		return post, nil
	} else {
		return post, errors.New("post is not found")
	}
}
