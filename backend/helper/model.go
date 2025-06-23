package helper

import (
	"github.com/fzndps/mini-social-media/backend/models/domain"
	"github.com/fzndps/mini-social-media/backend/models/web"
)

func ToUserResponse(user domain.User) web.UserRegisterResponse {
	return web.UserRegisterResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}
}

func ToUserLoginResponse(user domain.User, token string) web.UserLoginResponse {
	return web.UserLoginResponse{
		TokenType:   "Bearer ",
		AccessToken: token,
		ExpiresIn:   3600,
		User: web.UserRegisterResponse{
			Id:       user.Id,
			Username: user.Username,
			Email:    user.Email,
		},
	}
}

func ToCreatePostResponse(post domain.Post) web.PostCreateResponse {
	return web.PostCreateResponse{
		Id:        post.Id,
		UserId:    post.UserId,
		Content:   post.Content,
		ImageURL:  post.ImageURL,
		CreatedAt: post.CreatedAt,
	}
}

func ToPostResponses(posts []domain.Post) []web.FindAllPostResponses {
	var postResponses []web.FindAllPostResponses
	for _, post := range posts {
		postResponses = append(postResponses, web.FindAllPostResponses{
			Id:           post.Id,
			Content:      post.Content,
			ImageURL:     post.ImageURL,
			CreatedAt:    post.CreatedAt,
			CommentCount: post.CommentCount,
			User: web.UserResponse{
				Id:       post.User.Id,
				Username: post.User.Username,
			},
		})
	}
	return postResponses
}

func ToUserPostResponses(user domain.UserwithPost) web.UserPostResponse {
	var posts []web.UserPostProfile

	for _, post := range user.Posts {
		posts = append(posts, web.UserPostProfile{
			PostId:    post.PostId,
			Content:   post.Content,
			ImageURL:  post.ImageURL,
			CreatedAt: post.CreatedAt,
		})
	}

	return web.UserPostResponse{
		Id:       user.Id,
		Username: user.Username,
		Posts:    posts,
	}
}

func ToCommentResponse(comment domain.CommentWitUser) web.CommentResponse {
	return web.CommentResponse{
		Id:      comment.Id,
		Content: comment.Content,
		// PostId:    comment.PostId,
		CreatedAt: comment.CreatedAt,
		User: web.UserInfo{
			// Id:       comment.User.Id,
			Username: comment.User.Username,
		},
	}
}

func ToCommentResponses(comments []domain.CommentWitUser) []web.CommentResponse {
	var commentresponses []web.CommentResponse

	for _, comment := range comments {
		commentresponses = append(commentresponses, web.CommentResponse{
			Id:      comment.Id,
			Content: comment.Content,
			// PostId:    comment.PostId,
			CreatedAt: comment.CreatedAt,
			User: web.UserInfo{
				// Id:       comment.User.Id,
				Username: comment.User.Username,
			},
		})
	}

	return commentresponses
}
