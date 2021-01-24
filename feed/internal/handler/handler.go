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
}

// HelloPubSub consumes a Pub/Sub message.
func HelloPubSub(ctx context.Context, m PubSubMessage) error {
	message := post.PostTopicMessage{}
	err := json.Unmarshal(m.Data, &message)
	if err != nil{
		log.Println(err)
		return err
	}

	log.Println(message)

	return nil
}

































