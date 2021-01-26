package domain

import (
	"errors"
	"regexp"
)

type MediaType struct {
	MediaSrc *string `json:"mediaSrc";firestore:"mediaSrc,omitempty"`
	IsStill *bool `json:"isStill";firestore:"isStill,omitempty"`
}

// The most inner impl
type Tweet struct {
	Id       * string    `json:"id";firestore:"id,omitempty"`
	Text     * string    `json:"text";firestore:"text,omitempty"`
	Media    * MediaType `json:"media";firestore:"media,omitempty"`
	Hashtags * []string  `json:"hashtags";firestore:"hashtags,omitempty"`
}


func NewTweet(NewID *  string, Text * string, Media * MediaType) (* Tweet, error){

	if len(* Text) > 250 || len(* Text) == 0 {
		return nil, errors.New("Text length not valid")
	}

	r := regexp.MustCompile(`#[a-zA-Z0-9@]+`)
	Hashtags := r.FindAllString(*Text, -1)

	return &Tweet{NewID,Text, Media, &Hashtags}, nil
}







