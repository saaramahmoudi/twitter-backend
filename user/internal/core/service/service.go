package service

import (
	"context"
	"errors"
	"github.com/saaramahmoudi/twitter-backend/user/internal/core/domain"
	"github.com/saaramahmoudi/twitter-backend/user/internal/core/ports"
)

// The second most inner impl
// TODO good note on the doc on how we handled messaging between adapter and this port and the other port of auth
type UserService struct {
	Repo ports.UserRepository
	Auth ports.UserAuthenticator
}
func (u UserService) Get(ctx context.Context, email *  string) (* domain.User, error){
	return u.Repo.Get(email)
}

//This is the example of why we have a repo and a domain service at the same time
func (u UserService) UpdateId(ctx context.Context, id * string) (* domain.User, error){
	if !u.Auth.IsLoggedIn(ctx) {
		return nil, errors.New("User should be logged in for id update")
	}

	email, err := u.Auth.GetEmail(ctx)
	if err != nil {
		return nil, err
	}
	user, err := u.Get(ctx, email)

	if err != nil{
		return user, err
	}
	user, err = domain.NewUser(user.Name, user.Email, id, user.ImageSrc)
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

// Another example of why we need domain service, for things related to buss logic but outside of pure domain and not cross domain
func (u UserService) Create(ctx context.Context) (* domain.User, error) {

	if !u.Auth.IsLoggedIn(ctx) {
		return nil, errors.New("User should be logged in for id update")
	}

	authEmail, err := u.Auth.GetEmail(ctx)
	if err != nil {
		return nil, err
	}

	if u.Repo.EmailExists(authEmail) {
		return nil, errors.New("User already has an account")
	}
	name := ""
	imageSrc := ""
	user, err := domain.NewUser(&name, authEmail, authEmail, &imageSrc)
	if err != nil {
		return nil, err
	}
	user, err = u.Repo.Save(user)
	return user, err

}


func (u UserService) EmailExists(ctx context.Context) (* bool, error) {
	email, err := u.Auth.GetEmail(ctx)
	if err != nil {
		return nil, err
	}
	res := u.Repo.EmailExists(email)
	return &res, nil
}



