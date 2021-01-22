package post


// Package p contains a Pub/Sub Cloud Function.

import (
"context"
"log"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type InputAttribute struct {
	Email * string `json:"email"`
}

type PubSubMessage struct {
	Data []byte `json:"data"`
	Attributes * InputAttribute `json:"attributes"`
}

// HelloPubSub consumes a Pub/Sub message.
func HelloPubSub(ctx context.Context, m PubSubMessage) error {
	if m.Attributes.Email == nil {
		log.Fatal("Could not find email")
	}
	log.Println(* m.Attributes.Email)
	return nil
}
