package main

import (
	"github.com/saaramahmoudi/twitter-backend/tweet/pkg/core/domain"
	"github.com/saaramahmoudi/twitter-backend/tweet/pkg/handler"
	"log"
)

func main(){

	// Test creating
	text := "Something #new"
	tweet, err := handler.TweetApi.Create(&text, nil)

	if err != nil {
		log.Fatal(err)
	}

	add := "something"
	still := true
	tweet.Media = &domain.MediaType{MediaSrc: &add,  IsStill: &still}

	tweet, err = handler.TweetApi.Update(tweet.ID, tweet)

	if err != nil {
		log.Fatal(err)
	}

	err = handler.TweetApi.Delete(tweet)

	if err != nil {
		log.Fatal(err)
	}

}


