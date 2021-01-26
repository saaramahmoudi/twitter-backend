package domain

import (
	"time"
	"github.com/saaramahmoudi/twitter-backend/post"
	"github.com/saaramahmoudi/twitter-backend/user"
	)
type EventEnum = string

type PubSubMessage struct {
	Data []byte `json:"data"`
	Attributes map[string]string `json:"attributes"`
}
const (
	PostPublish = "PostPublish"
	PostLike = "PostLiked"
	PostRetweet = "PostRetweet"
	UserFollow = "UserFollow"
	PostRetweeted = "PostRetweeted"
)
type Event struct {
	EventType EventEnum `json:"eventType"`
	MadeByUserId * string `json:"madeByUserId"`
	MadeAt * time.Time `json:"madeAt"`
	IsReversal bool `json:"isReversal"`
}

type PostEvent struct {
	Post * post.Post `json:"post"`
	Event
}


type UserEvent struct {
	UserMain * user.User `json:"userMain"`
	UserSecondary * user.User `json:"userSecondary"`
	Event
}



























