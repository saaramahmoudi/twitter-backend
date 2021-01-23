package user

import (
	"github.com/saaramahmoudi/twitter-backend/user/internal/core/domain"
	"github.com/saaramahmoudi/twitter-backend/user/internal/core/ports"
)

type UserApi struct {
	ports.UserService
}
type UserAuth struct {
	ports.HttpUserAuthenticator
}

type User = domain.User
var ApiAuth = UserAuth{HttpUserAuthenticator: authHandler}
var Api = UserApi{UserService: userService}





