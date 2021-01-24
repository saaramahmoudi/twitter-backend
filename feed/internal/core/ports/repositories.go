package ports

import (
	"github.com/saaramahmoudi/twitter-backend/feed/internal/core/domain"
)

type FeedRepository interface {
	SaveFeed (feed  domain.FeedProvider) (domain.FeedProvider, error)
	GetFeed (feed  domain.FeedProvider, id * string) (domain.FeedProvider, error)
}



