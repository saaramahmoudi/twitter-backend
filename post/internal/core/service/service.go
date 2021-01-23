package service

import (
	"context"
	"github.com/saaramahmoudi/twitter-backend/post/internal/core/domain"
	"github.com/saaramahmoudi/twitter-backend/post/internal/core/ports"
	"github.com/saaramahmoudi/twitter-backend/tweet"
	"github.com/saaramahmoudi/twitter-backend/user"
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
func (u PostService) Create(ctx context.Context, Text * string, MediaType *  tweet.MediaType) (* domain.Post, error) {
	tweetInstance, err := tweet.Api.Create(Text, MediaType)
	if err != nil {
		return nil, err
	}

	email, err := user.ApiAuth.GetEmail(ctx)
	if err != nil{
		return nil, err
	}

	userInstance, err := user.Api.GetByEmail(ctx, email)
	if err != nil{
		return nil, err
	}

	id, err :=  u.Repo.GetNewId()
	if err != nil{
		return nil, err
	}

	post, err := domain.NewPost(id, userInstance.Id, tweetInstance.ID, []string{}, []string{})
	if err != nil{
		return nil, err
	}

	post, err = u.Repo.Save(post)

	return post, err
}

