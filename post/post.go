package post

import (
	"github.com/saaramahmoudi/twitter-backend/post/internal/core/domain"
)

type Post = domain.Post
type PostEvent = domain.PostEvent
type EventEnum = domain.EventEnum
const PostLiked = domain.PostLiked
const PostRetweeted = domain.PostRetweeted
const PostCreated = domain.PostCreated
