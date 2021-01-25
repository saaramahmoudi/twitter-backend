package main

import (
	tweet "github.com/saaramahmoudi/twitter-backend/tweet"
	"log"
)

func main(){

	// Test creating
	text := "#wasd1 something #wasd asd #wasd2"
	tweetInstance, err := tweet.Api.Create(&text, nil)

	if err != nil {
		log.Fatal(err)
	}

	add := "something"
	still := true
	tweetInstance.Media = &tweet.MediaType{MediaSrc: &add,  IsStill: &still}

	tweetInstance, err = tweet.Api.Update(tweetInstance.Id, tweetInstance)

	if err != nil {
		log.Fatal(err)
	}

	//err = tweet.Api.Delete(tweetInstance)

	if err != nil {
		log.Fatal(err)
	}

}


