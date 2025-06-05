package services

import (
	"context"

	"github.com/fzndps/mini-social-media/backend/models/web"
)

type PostService interface {
	Create(ctx context.Context, request web.PostCreateRequest) web.PostCreateResponse
	FindAll(ctx context.Context) []web.FindAllPostResponses
}
