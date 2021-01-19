package repositories

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	firebase "firebase.google.com/go"
	"github.com/saaramahmoudi/twitter-backend/user/internal/core/domain"
	"google.golang.org/api/iterator"
	"log"
)

type UserFirestore struct{

}

var client *firestore.Client
var ctx = context.Background()
var app *firebase.App

func (repo UserFirestore) Get(email string) (*domain.User, error){
	doc, err := client.Collection("UserProfile").Doc(email).Get(ctx)
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

func (repo UserFirestore) GetUserFromId(id string) (* domain.User, error){
	iter := client.Collection("UserProfile").Where("id", "==", id).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		user := domain.User{}
		err = doc.DataTo(&user)
		if err != nil {
			continue
		}
		//This is because the ids are also unique to users, TODO remove this dep
		return &user, nil
	}
	return nil, errors.New("Couldn't find the user")
}
func (repo UserFirestore) UpdateUser(user * domain.User) (* domain.User, error) {
	_, err := client.Collection("UserProfile").Doc(user.Email).Set(ctx, user)
	if err != nil{
		return nil, err
	}
	return user, nil
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









