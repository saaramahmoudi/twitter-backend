package repositories

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	firebase "firebase.google.com/go"
	"github.com/saaramahmoudi/twitter-backend/feed/internal/core/domain"
	"log"
	"reflect"
)

type FeedFirestore struct{

}

var client *firestore.Client
var ctx = context.Background()
var app *firebase.App
// This is used because it seems that the gcp firestore library does not use json tags correctly for creating documents
func turnStructToMap(input interface{}) (map[string]interface{}, error) {
	bytes, err := json.Marshal(&input)
	if err != nil {
		return nil, err
	}
	var res map[string]interface{}
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}



func (f FeedFirestore) SaveFeed (feed domain.FeedProvider) (domain.FeedProvider, error) {
	mapJson, err := turnStructToMap(feed.GetFeed())
	if err != nil {
		return nil, err
	}
	_, err = client.Collection(reflect.TypeOf(feed).Name()).Doc(*feed.GetFeed().Id).Set(ctx, mapJson)
	if err != nil {
		return nil, err
	}
	return feed, nil
}
func (f FeedFirestore) GetFeed (feed domain.FeedProvider, id * string) (domain.FeedProvider, error) {
	doc, err := client.Collection(reflect.TypeOf(feed).Name()).Doc(*id).Get(ctx)

	if err != nil {
		return nil, err
	}

	err = doc.DataTo(feed.GetFeed())

	if err != nil {
		return nil, err
	}

	return feed, nil
}

func init(){
	//TODO clean up inits
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









