package main

import (
	"github.com/saaramahmoudi/twitter-backend/tweet/pkg/core/domain"
	"github.com/saaramahmoudi/twitter-backend/tweet/pkg/tweet"
	"log"
)

func main(){

	// Test creating
	text := "Something #new"
	tweetInstance, err := tweet.TweetApi.Create(&text, nil)

	if err != nil {
		log.Fatal(err)
	}

	add := "something"
	still := true
	tweetInstance.Media = &domain.MediaType{MediaSrc: &add,  IsStill: &still}

	tweetInstance, err = tweet.TweetApi.Update(tweetInstance.ID, tweetInstance)

	if err != nil {
		log.Fatal(err)
	}

	err = tweet.TweetApi.Delete(tweetInstance)

	if err != nil {
		log.Fatal(err)
	}

}


