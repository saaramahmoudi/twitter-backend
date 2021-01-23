package ports

import (
	"context"
	"github.com/saaramahmoudi/twitter-backend/user/internal/core/domain"
)

//The reason service exists, is separation of concern because the repo should not be concerned with the buss logic even if there is little of it
// This could have also gon into the domain cause it is a pure domain service
type UserService interface {
	GetByEmail(ctx context.Context, email * string) (* domain.User, error)
	UpdateTag(ctx context.Context, tag * string) (* domain.User, error)
	Create(ctx context.Context) (* domain.User, error)
	EmailExists(ctx context.Context) (* bool, error)
}











