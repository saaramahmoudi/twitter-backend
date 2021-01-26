package handler

import (
	"context"
	"encoding/json"
	"github.com/saaramahmoudi/twitter-backend/user"
	"github.com/saaramahmoudi/twitter-backend/post"
	"log"
)


type MessageHandler struct {
	user.UserApi
}


type PubSubMessage struct {
	Data []byte `json:"data"`
	Attributes map[string]string `json:"attributes"`
}

// HelloPubSub consumes a Pub/Sub message.
func (mh MessageHandler) Handle(ctx context.Context, m PubSubMessage) error {
	message := post.PostTopicMessage{}
	err := json.Unmarshal(m.Data, &message)
	if err != nil{
		log.Println(err)
		return err
	}

	if message.EventType ==  post.PostPublished {
		log.Println("Post published")
	} else if message.EventType == post.PostRetweeted {
		log.Println("Post retweeted")
	} else if message.EventType == post.PostLiked {
		log.Println("Post liked")
	}

	return nil
}

































