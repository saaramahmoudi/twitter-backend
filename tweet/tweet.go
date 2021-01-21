package tweet

import (
	"github.com/saaramahmoudi/twitter-backend/tweet/internal/core/domain"
	"github.com/saaramahmoudi/twitter-backend/tweet/internal/core/ports"
	"github.com/saaramahmoudi/twitter-backend/tweet/internal/core/service"
	"github.com/saaramahmoudi/twitter-backend/tweet/internal/repositories"
)

type tweetApi struct {
	ports.TweetService
}

var NewTweet = domain.NewTweet
type MediaType domain.MediaType
type Tweet domain.Tweet
var Api = tweetApi{service.TweetService{Repo :repositories.TweetFirestore{}}}






















