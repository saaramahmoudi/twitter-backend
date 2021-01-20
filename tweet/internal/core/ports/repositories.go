package ports

import "github.com/saaramahmoudi/twitter-backend/tweet/internal/core/domain"

type TweetRepository interface {
	//Save(Text string, Media domain.MediaType) (*domain.Tweet, error)
	GetTweet(id string) (* domain.Tweet, error)
	UpdateTweet(user * domain.Tweet) (* domain.Tweet, error)
}



