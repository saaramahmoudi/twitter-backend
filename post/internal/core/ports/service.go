package ports

import (
	"context"
	"github.com/saaramahmoudi/twitter-backend/post/internal/core/domain"
)

//The reason service exists, is separation of concern because the repo should not be concerned with the buss logic even if there is little of it
// This could have also gon into the domain cause it is a pure domain service
type PostService interface {
	Get(ctx context.Context, id * string) (* domain.Post, error)
	Create(ctx context.Context) (* domain.Post, error)
}











