package ports

import "github.com/saaramahmoudi/twitter-backend/user/internal/core/domain"

type UserRepository interface {
	Get(email * string) (*domain.User, error)
	GetUserFromTag(tag * string) (* domain.User, error)
	UpdateUser(user * domain.User) (* domain.User, error)
	EmailExists(email * string) bool
	Save(user * domain.User) (* domain.User, error)
}



