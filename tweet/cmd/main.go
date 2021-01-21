package main

import (
	tweet2 "github.com/saaramahmoudi/twitter-backend/tweet"
	"github.com/saaramahmoudi/twitter-backend/tweet/internal/core/domain"
	"log"
)

func main(){

	// Test creating
	text := "Something #new"
	tweetInstance, err := tweet2.Api.Create(&text, nil)

	if err != nil {
		log.Fatal(err)
	}

	add := "something"
	still := true
	tweetInstance.Media = &domain.MediaType{MediaSrc: &add,  IsStill: &still}

	tweetInstance, err = tweet2.Api.Update(tweetInstance.ID, tweetInstance)

	if err != nil {
		log.Fatal(err)
	}

	err = tweet2.Api.Delete(tweetInstance)

	if err != nil {
		log.Fatal(err)
	}

}


