package ports

import "github.com/saaramahmoudi/twitter-backend/post/internal/core/domain"
const NoEventFound = "No event found"
type PostRepository interface {
	Get(id * string) (*domain.Post, error)
	Save(post * domain.Post) (* domain.Post, error)
	SaveOrDeleteEvent(post * domain.PostEvent) error
	GetNewId() (*string, error)
	GetSaveTransaction(id * string, operation func (* domain.Post) (* domain.Post, error)) (* domain.Post, error)
}



