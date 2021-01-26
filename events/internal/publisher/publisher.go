package publisher

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"errors"
	"github.com/saaramahmoudi/twitter-backend/events/internal/core/domain"
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

func (ep * EventPublisher) PublishPostEvent(event * domain.PostEvent) (* string, error) {
	t := client.Topic("EventFeed")
	bytes, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	//TODO the only dep here is the "type" remove this using reflection in future
	result := t.Publish(ctx, &pubsub.Message{
		Data: bytes,
		Attributes: map[string]string{
			"type":   "Post",
		},
	})


	id, err := result.Get(ctx)


	return &id, err
}







func (ep * EventPublisher) PublishUserEvent(event * domain.UserEvent) (* string, error) {

	t := client.Topic("EventFeed")

	bytes, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	result := t.Publish(ctx, &pubsub.Message{
		Data: bytes,
		Attributes: map[string]string{
			"type":   "User",
		},
	})


	id, err := result.Get(ctx)


	return &id, err
}
func (ep * EventPublisher) CheckIfPostEvent(message * domain.PubSubMessage) (* domain.PostEvent, error){
	if message.Attributes["type"] == "Post" {
		event := &domain.PostEvent{}
		err := json.Unmarshal(message.Data, event)
		return event, err
	}else {
		return nil, errors.New("Not Post Event")
	}
}
func (ep * EventPublisher) CheckIfUserEvent(message * domain.PubSubMessage) (* domain.UserEvent, error){
	if message.Attributes["type"] == "User" {
		event := &domain.UserEvent{}
		err := json.Unmarshal(message.Data, event)
		return event, err

	}else {
		return nil, errors.New("Not User Event")
	}
}













