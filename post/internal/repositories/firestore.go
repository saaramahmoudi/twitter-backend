package repositories

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	firebase "firebase.google.com/go"
	"github.com/saaramahmoudi/twitter-backend/post/internal/core/domain"
	"log"
)

type PostFirestore struct{

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
func (repo PostFirestore) Get(id * string) (*domain.Post, error){
	doc, err := client.Collection("UserProfile").Doc(*email).Get(ctx)
	if err != nil {
		return nil, err
	}

	res := domain.User{}
	errDataTransfer := doc.DataTo(&res)

	if errDataTransfer != nil {
		return nil, errDataTransfer
	}
	// This is because we don't keep the email on the firestore to avoid repetition TODO check if it's a good design
	res.Email = email
	return &res, nil
}

//TODO check if we need to merge update and save
func (repo PostFirestore) Save(post * domain.Post) (* domain.Post, error){
	mapUser, err := turnStructToMap(user)
	if err != nil{
		return nil, err
	}
	_, err = client.Collection("UserProfile").Doc(* user.Email).Set(ctx, mapUser)
	if err != nil{
		return nil, err
	}
	return user, nil
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









