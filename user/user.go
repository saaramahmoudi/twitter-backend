package user

import (
	"github.com/saaramahmoudi/twitter-backend/user/internal/core/ports"
)

type UserApi struct {
	ports.UserService
}
type UserAuth struct {
	ports.UserAuthenticator
}

var ApiAuth = UserAuth{UserAuthenticator: authHandler}
var Api = UserApi{UserService: userService}





