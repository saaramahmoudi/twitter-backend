package publisher

import (
	"cloud.google.com/go/pubsub"
	"context"
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
func (ep * EventPublisher) Publish(event * domain.PostEvent) (* string, error) {

	t := client.Topic(event.EventType)

	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(event.EventType),
		Attributes: map[string]string{
			"id":   *event.Post.Id,
		},
	})


	id, err := result.Get(ctx)


	return &id, err
}













