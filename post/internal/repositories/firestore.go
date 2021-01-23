package repositories

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"errors"
	firebase "firebase.google.com/go"
	"github.com/saaramahmoudi/twitter-backend/post/internal/core/domain"
	"log"
)

type PostFirestore struct{

}

var client *firestore.Client
var ctx = context.Background()
var app *firebase.App
const CollectionAddress = "Posts"
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
func (repo PostFirestore) Get(id * string) (*domain.Post, error){
	doc, err := client.Collection(CollectionAddress).Doc(*id).Get(ctx)
	if err != nil {
		return nil, err
	}

	res := domain.Post{}
	errDataTransfer := doc.DataTo(&res)

	if errDataTransfer != nil {
		return nil, errDataTransfer
	}
	return &res, nil
}

func (repo PostFirestore) GetNewId() (*string, error) {
	ref := client.Collection(CollectionAddress).NewDoc()
	if ref.ID == ""{
		return nil, errors.New("Could not create a new post")
	}
	return &ref.ID, nil
}

//TODO check if we need to merge update and save
func (repo PostFirestore) Save(post * domain.Post) (* domain.Post, error){
	mapUser, err := turnStructToMap(post)
	if err != nil{
		return nil, err
	}
	_, err = client.Collection(CollectionAddress).Doc(* post.Id).Set(ctx, mapUser)
	if err != nil{
		return nil, err
	}
	return post, nil
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









