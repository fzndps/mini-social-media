package repository

import (
	"context"
	"database/sql"

	"github.com/fzndps/mini-social-media/backend/helper"
	"github.com/fzndps/mini-social-media/backend/models/domain"
)

type UserPostRepositoryImpl struct {
	DB *sql.DB
}

func NewUserPostRepository(DB *sql.DB) UserPostRepository {
	return &UserPostRepositoryImpl{
		DB: DB,
	}
}

func (repository *UserPostRepositoryImpl) UserPostProfile(ctx context.Context, userId int) domain.UserwithPost {
	SQL := `SELECT 
    				u.id AS user_id, u.username,
    				p.id AS post_id, p.content, p.image_url, p.created_at,
					(SELECT COUNT(*) FROM comments c WHERE c.post_id = p.id) AS comment_count
					FROM users u
					LEFT JOIN posts p ON u.id = p.user_id
					WHERE u.id = ?;
					`
	rows, err := repository.DB.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	var user domain.UserwithPost
	user.Posts = []domain.UserProfilePost{}

	for rows.Next() {
		var (
			post domain.UserProfilePost
		)

		err := rows.Scan(&user.Id, &user.Username, &post.PostId, &post.Content, &post.ImageURL, &post.CreatedAt, &post.CommentCount)
		helper.PanicIfError(err)

		user.Posts = append(user.Posts, post)
	}

	return user
}
