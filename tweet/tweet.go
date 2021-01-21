package tweet

import (
	"github.com/saaramahmoudi/twitter-backend/tweet/internal/core/ports"
	"github.com/saaramahmoudi/twitter-backend/tweet/internal/core/service"
	"github.com/saaramahmoudi/twitter-backend/tweet/internal/repositories"
	"github.com/saaramahmoudi/twitter-backend/tweet/pkg/core/domain"
)

type tweetApi struct {
	ports.TweetService
}

type MediaType domain.MediaType
type Tweet domain.Tweet
var Api = tweetApi{service.TweetService{Repo :repositories.TweetFirestore{}}}






















