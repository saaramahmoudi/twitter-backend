package user

import (
	"github.com/saaramahmoudi/twitter-backend/user/internal/core/ports"
)

type UserApi struct {
	ports.UserService
}

var Api = UserApi{UserService: userService}





