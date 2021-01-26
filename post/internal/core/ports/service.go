package ports

import (
	"context"
	"github.com/saaramahmoudi/twitter-backend/post/internal/core/domain"
	"github.com/saaramahmoudi/twitter-backend/tweet"
)

//The reason service exists, is separation of concern because the repo should not be concerned with the buss logic even if there is little of it
// This could have also gon into the domain cause it is a pure domain service
type PostService interface {
	Get(ctx context.Context, id * string) (* domain.Post, error)
	Create(ctx context.Context, Text * string, MediaType * tweet.MediaType) (* domain.PostEvent, error)
	ToggleLike(ctx context.Context, postId * string) (* domain.PostEvent, error)
	ToggleRetweet(ctx context.Context, postId * string) (* domain.PostEvent, error)
}











