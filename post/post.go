package post

import "github.com/saaramahmoudi/twitter-backend/post/internal/core/domain"

type Post = domain.Post
type PostEvent = domain.PostEvent
type EventEnum = domain.EventEnum
const PostPublished = domain.PostPublished
const PostLiked = domain.PostLiked
const PostRetweeted = domain.PostRetweeted


