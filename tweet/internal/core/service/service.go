package service

import (
	"github.com/saaramahmoudi/twitter-backend/tweet/internal/core/domain"
	"github.com/saaramahmoudi/twitter-backend/tweet/internal/core/ports"
)

// The second most inner impl

type TweetService struct {
	Repo ports.TweetRepository
}
func (t TweetService) Get(id * string) (*domain.Tweet, error){
	return t.Repo.GetTweet(id)
}

//This is the example of why we have a repo and a domain service at the same time
func (t TweetService) Update(id * string, tweet *domain.Tweet) (*domain.Tweet, error){
	oldTweet, err := t.Get(id)
	if err != nil{
		return nil, err
	}

	tweet.ID = oldTweet.ID
	if tweet.Text == nil {
		tweet.Text = oldTweet.Text
	}
	if tweet.Media == nil{
		tweet.Media = oldTweet.Media
	}

	tweet, err = domain.NewTweet(tweet.ID, tweet.Text, tweet.Media)

	if err != nil {
		return nil, err
	}

	tweet, err = t.Repo.UpdateTweet(tweet)
	if err != nil{
		return nil, err
	}

	return tweet, nil
}

func (t TweetService) Create(Text * string, Media *domain.MediaType) (*domain.Tweet, error){
	id, err := t.Repo.GetNewId()
	if err != nil{
		return nil, err
	}
	tweet, err := domain.NewTweet(id, Text, Media)
	if err != nil{
		return nil, err
	}
	tweet, err = t.Repo.Save(tweet)
	if err != nil {
		return nil, err
	}
	return tweet, nil
}

func (t TweetService)  Delete(tweet *domain.Tweet) error {
	return t.Repo.Delete(tweet)
}



