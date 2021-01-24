package publisher

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"github.com/saaramahmoudi/twitter-backend/post/internal/core/domain"
	"log"
)

type EventPublisher struct {

}

var client * pubsub.Client
var ctx context.Context

func init (){
	ctx = context.Background()
	var err error
	client, err = pubsub.NewClient(ctx, "twitter-1db2a")
	if err != nil {
		log.Fatal(err)
	}
}

type PostTopicMessage struct {
	PostId * string `json:"postId"`
	EventType domain.EventEnum `json:"eventType"`
	MadeByUserId * string `json:"madeByUserId"`
}

func (ep * EventPublisher) Publish(event * domain.PostEvent) (* string, error) {

	t := client.Topic("PostFeed")

	message := PostTopicMessage{
		PostId: event.Post.Id,
		EventType: event.EventType,
		MadeByUserId: event.MadeByUserId,
	}

	bytes, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	result := t.Publish(ctx, &pubsub.Message{
		Data: bytes,
	})


	id, err := result.Get(ctx)


	return &id, err
}













