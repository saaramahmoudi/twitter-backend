package service

import (
	"context"
	"errors"
	"github.com/saaramahmoudi/twitter-backend/post/internal/core/domain"
	"github.com/saaramahmoudi/twitter-backend/post/internal/core/ports"
)

// The second most inner impl
// TODO good note on the doc on how we handled messaging between adapter and this port and the other port of auth
type PostService struct {
	Repo ports.PostRepository
}
func (u PostService) Get(ctx context.Context, id *  string) (* domain.Post, error){
	return u.Repo.Get(id)
}

// Another example of why we need domain service, for things related to buss logic but outside of pure domain and not cross domain
func (u PostService) Create(ctx context.Context) (* domain.Post, error) {

}

