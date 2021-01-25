package repositories

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	firebase "firebase.google.com/go"
	"github.com/saaramahmoudi/twitter-backend/post/internal/core/domain"
	"github.com/saaramahmoudi/twitter-backend/utils"
	"log"
)

type PostFirestore struct{

}

var client *firestore.Client
var ctx = context.Background()
var app *firebase.App
const CollectionAddress = "Posts"
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
	mapUser, err := utils.TurnStructToMap(post)
	log.Println(mapUser)
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









