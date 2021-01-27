package service

import (
	"context"
	"github.com/saaramahmoudi/twitter-backend/post/internal/core/domain"
	"github.com/saaramahmoudi/twitter-backend/post/internal/core/ports"
	"github.com/saaramahmoudi/twitter-backend/tweet"
	"github.com/saaramahmoudi/twitter-backend/user"
	"time"
)

// The second most inner impl
// TODO good note on the doc on how we handled messaging between adapter and this port and the other port of auth
type PostService struct {
	Repo ports.PostRepository
	Auth user.UserAuth
}
func (u PostService) Get(ctx context.Context, id *  string) (* domain.Post, error){
	return u.Repo.Get(id)
}

// Another example of why we need domain service, for things related to buss logic but outside of pure domain and not cross domain
func (u PostService) Create(ctx context.Context, Text * string, MediaType *  tweet.MediaType) (* domain.Post, error) {
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


	currTime := time.Now().Unix()
	post, err := domain.NewPost(id, userInstance.Id, tweetInstance.Id, []string{}, []string{}, &currTime)
	if err != nil{
		return nil, err
	}
	_, err = u.Repo.Save(post)
	if err != nil {
		return nil, err
	}
	isReversal := false
	pe := &domain.PostEvent{PostId: post.Id, EventType: domain.PostCreated, MadeByUserId:  userInstance.Id, MadeAt: &currTime, IsReversal: &isReversal}
	err = u.Repo.SaveOrDeleteEvent(pe)
	return post, err
}

func (u PostService) ToggleLike(ctx context.Context, postId * string) (* domain.PostEvent, error) {
	email, err := u.Auth.GetEmail(ctx)
	if err != nil{
		return nil, err
	}

	userInstance, err := user.Api.GetByEmail(ctx, email)
	if err != nil{
		return nil, err
	}
	var pe  * domain.PostEvent

	operation := func (* domain.Post) (* domain.Post, error){
		post, err := u.Get(ctx, postId)
		if err != nil{
			return nil, err
		}

		currTime := time.Now().Unix()
		if err != nil {
			return nil, err
		}

		pe, err = post.ToggleLikePost(userInstance.Id, &currTime)
		if err != nil{
			return nil, err
		}
		err = u.Repo.SaveOrDeleteEvent(pe)
		//This is just in case we manually delete an event
		if err != nil &&  err.Error() == ports.NoEventFound {
			err = nil
		}
		return post, nil
	}

	_, err = u.Repo.GetSaveTransaction(postId, operation)
	return pe, err
}


func (u PostService) ToggleRetweet(ctx context.Context, postId * string) (* domain.PostEvent, error) {
	email, err := u.Auth.GetEmail(ctx)
	if err != nil{
		return nil, err
	}

	userInstance, err := user.Api.GetByEmail(ctx, email)
	if err != nil{
		return nil, err
	}
	var pe  * domain.PostEvent
	operation := func (* domain.Post) (* domain.Post, error){
		post, err := u.Get(ctx, postId)
		if err != nil{
			return nil, err
		}

		currTime := time.Now().Unix()
		if err != nil {
			return nil, err
		}
		pe, err = post.ToggleRetweetPost(userInstance.Id, &currTime)
		if err != nil{
			return nil, err
		}
		err = u.Repo.SaveOrDeleteEvent(pe)
		//This is just in case we manually delete an event
		if err != nil && err.Error() == ports.NoEventFound {
			err = nil
		}

		return post, err
	}

	_, err = u.Repo.GetSaveTransaction(postId, operation)
	return pe, err
}

