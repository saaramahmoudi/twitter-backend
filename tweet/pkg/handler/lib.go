package handler

import (
	"github.com/saaramahmoudi/twitter-backend/tweet/internal/core/ports"
	"github.com/saaramahmoudi/twitter-backend/tweet/internal/core/service"
	"github.com/saaramahmoudi/twitter-backend/tweet/internal/repositories"
)

type tweetApi struct {
	ports.TweetService
}

var TweetApi = tweetApi{service.TweetService{Repo :repositories.TweetFirestore{}}}






















