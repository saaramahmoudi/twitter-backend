package repositories

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/saaramahmoudi/twitter-backend/tweet/internal/core/domain"

	//"errors"
	firebase "firebase.google.com/go"
	"log"
)

const CollectionAddress = "Tweet"

type TweetFirestore struct{

}

var client *firestore.Client
var ctx = context.Background()
var app *firebase.App

//func (tf TweetFirestore) Save(Text string, Media domain.MediaType) (*domain.Tweet, error){
//
//}

func (tf TweetFirestore) GetTweet(id string) (* domain.Tweet, error) {
	res := &domain.Tweet{}
	doc, err := client.Collection(CollectionAddress).Doc(id).Get(ctx)
	if err != nil{
		log.Println(err)
		return nil, err
	}

	if err = doc.DataTo(res); err != nil{
		log.Println(err)
		return nil, err
	}

	return res, nil
}
func (tf TweetFirestore) UpdateTweet(tweet * domain.Tweet) (* domain.Tweet, error) {
	_, err := client.Collection(CollectionAddress).Doc(*tweet.ID).Set(ctx, tweet)
	if err != nil{
		log.Println(err)
		return nil, err
	}
	return tweet, nil
}


func init(){
	var err error
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	app, err = firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}









