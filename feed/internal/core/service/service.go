package service

import (
	"context"
	"github.com/saaramahmoudi/twitter-backend/feed/internal/core/domain"
	"github.com/saaramahmoudi/twitter-backend/feed/internal/core/ports"
)

// The second most inner impl
// TODO good note on the doc on how we handled messaging between adapter and this port and the other port of auth
type FeedService struct {
	Repo ports.FeedRepository
}



func (f FeedService) SaveFeed(ctx context.Context, id * string, feed domain.FeedProvider) (domain.FeedProvider, error){
	return f.Repo.SaveFeed(feed)
}
func (f FeedService) GetFeed(ctx context.Context, id * string, feed  domain.FeedProvider) (domain.FeedProvider, error){
	return f.Repo.GetFeed(feed, id)
}


