package events

import (
	"github.com/saaramahmoudi/twitter-backend/events/internal/core/domain"
	"github.com/saaramahmoudi/twitter-backend/events/internal/core/ports"
	"github.com/saaramahmoudi/twitter-backend/events/internal/publisher"
)

type EventPublisher = ports.EventPublisher
var PublisherHandler ports.EventPublisher = &publisher.EventPublisher{}


var PostPublish = domain.PostPublish
var PostLike = domain.PostLike
var PostRetweet = domain.PostRetweet
var UserFollow = domain.UserFollow
var PostRetweeted = domain.PostRetweeted











