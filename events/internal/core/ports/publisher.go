package ports

import "github.com/saaramahmoudi/twitter-backend/events/internal/core/domain"

type EventPublisher interface {
	PublishUserEvent(event * domain.UserEvent) (* string, error)
	PublishPostEvent(event * domain.PostEvent) (* string, error)
	CheckIfPostEvent(message * domain.PubSubMessage) (* domain.PostEvent, error)
	CheckIfUserEvent(message * domain.PubSubMessage) (* domain.UserEvent, error)
}




