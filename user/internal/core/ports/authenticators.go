package ports

import (
	"context"
)

type UserAuthenticator interface {
	GetEmail(ctx context.Context) (* string, error)
	GetUid(ctx context.Context) (* string, error)
	IsLoggedIn(ctx context.Context) bool
}

//TODO check if the layer level of this interface is correct
type HttpUserAuthenticator interface {
	UserAuthenticator
	//TODO expects a string on the return of the dict, make this decoupled
	GetAuthHeaderKey() string
}















