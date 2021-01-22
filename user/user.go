package user

import (
	"github.com/saaramahmoudi/twitter-backend/user/internal/core/ports"
)

type UserApi struct {
	ports.UserService
}

//var NewTweet = domain.NewTweet
//type MediaType = domain.MediaType
//type Tweet = domain.Tweet
//var Api = UserApi{UserService: service.UserService{Repo: repositories.UserFirestore{}, Auth: authenticators.FirebaseAuthenticator{}}, Us: authenticators.FirebaseAuthenticator{}











