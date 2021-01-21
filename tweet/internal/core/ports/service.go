package ports

import (
	"github.com/saaramahmoudi/twitter-backend/tweet/internal/core/domain"
)

//The reason service exists, is separation of concern because the repo should not be concerned with the buss logic even if there is little of it
// This could have also gon into the domain cause it is a pure domain service
type TweetService interface {
	Get(id * string) (*domain.Tweet, error)
	Update(id * string, tweet *domain.Tweet) (*domain.Tweet, error)
	Create(Text * string, Media *domain.MediaType) (*domain.Tweet, error)
	Delete(tweet *domain.Tweet) error
}











