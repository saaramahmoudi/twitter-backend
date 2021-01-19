package service

import (
	"errors"
	"github.com/saaramahmoudi/twitter-backend/user/internal/core/domain"
	"github.com/saaramahmoudi/twitter-backend/user/internal/core/ports"
)

// The second most inner impl

type UserService struct {
	Repo ports.UserRepository
}
func (u UserService) Get(email string) (* domain.User, error){
	return u.Repo.Get(email)
}

//This is the example of why we have a repo and a domain service at the same time
func (u UserService) UpdateId(email string, id string) (* domain.User, error){
	user, err := u.Get(email)
	if err != nil{
		return user, err
	}

	userWithId, err := u.Repo.GetUserFromId(id)
	if userWithId != nil{
		return nil, errors.New("User already exits")
	}

	user.Id = id
	user, err = u.Repo.UpdateUser(user)

	return user, err

}






