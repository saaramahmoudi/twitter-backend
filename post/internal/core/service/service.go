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
	Auth user.UserAuth
	Publisher ports.EventPublisher
}
func (u PostService) Get(ctx context.Context, id *  string) (* domain.Post, error){
	return u.Repo.Get(id)
}

// Another example of why we need domain service, for things related to buss logic but outside of pure domain and not cross domain
func (u PostService) Create(ctx context.Context, Text * string, MediaType *  tweet.MediaType) (* domain.PostEvent, error) {
	email, err := u.Auth.GetEmail(ctx)
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

	tweetInstance, err := tweet.Api.Create(Text, MediaType)
	if err != nil {
		return nil, err
	}


	postEvent, err := domain.NewPost(id, userInstance.Id, tweetInstance.Id, []string{}, []string{})
	if err != nil{
		return nil, err
	}

	_, err = u.Repo.Save(postEvent.Post)
	if err != nil{
		return nil, err
	}

	_, err = u.Publisher.Publish(postEvent)
	if err != nil{
		return nil, err
	}

	return postEvent, err
}

func (u PostService) LikePost(ctx context.Context, postId * string) (* domain.PostEvent, error) {
	email, err := u.Auth.GetEmail(ctx)
	if err != nil{
		return nil, err
	}

	userInstance, err := user.Api.GetByEmail(ctx, email)
	if err != nil{
		return nil, err
	}
	post, err := u.Get(ctx, postId)
	if err != nil{
		return nil, err
	}
	pe, err := post.LikePost(userInstance.Id)
	if err != nil {
		return nil, err
	}
	_, err = u.Publisher.Publish(pe)
	if err != nil {
		return nil, err
	}
	return pe,nil
}


func (u PostService) RetweetPost(ctx context.Context, postId * string) (* domain.PostEvent, error) {
	email, err := u.Auth.GetEmail(ctx)
	if err != nil{
		return nil, err
	}

	userInstance, err := user.Api.GetByEmail(ctx, email)
	if err != nil{
		return nil, err
	}
	post, err := u.Get(ctx, postId)
	if err != nil{
		return nil, err
	}


	pe, err := post.RetweetPost(userInstance.Id)
	if err != nil {
		return nil, err
	}
	_, err = u.Publisher.Publish(pe)
	if err != nil {
		return nil, err
	}
	return pe,nil
}

