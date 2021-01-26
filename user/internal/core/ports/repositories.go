package ports

import "github.com/saaramahmoudi/twitter-backend/user/internal/core/domain"

type UserRepository interface {
	GetByEmail(email * string) (*domain.User, error)
	GetById(id * string) (*domain.User, error)
	GetUserFromTag(tag * string) (* domain.User, error)
	UpdateUser(user * domain.User) (* domain.User, error)
	EmailExists(email * string) (* bool, error)
	Save(user * domain.User) (* domain.User, error)
	GetSaveTransactionTwoUsers(id1 * string, id2 * string, operation func (user1 * domain.User, user2 * domain.User) (* domain.User, * domain.User,  error)) (* domain.User, * domain.User, error)
}



