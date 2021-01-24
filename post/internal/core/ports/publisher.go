package ports

import "github.com/saaramahmoudi/twitter-backend/post/internal/core/domain"

type EventPublisher interface {
	Publish(event * domain.PostEvent) (* string, error)
}













