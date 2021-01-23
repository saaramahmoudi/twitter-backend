package ports

import "github.com/saaramahmoudi/twitter-backend/post/internal/core/domain"

type PostRepository interface {
	Get(id * string) (*domain.Post, error)
	Save(post * domain.Post) (* domain.Post, error)
}



