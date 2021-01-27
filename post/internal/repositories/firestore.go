package repositories

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	firebase "firebase.google.com/go"
	"github.com/saaramahmoudi/twitter-backend/post/internal/core/domain"
	"github.com/saaramahmoudi/twitter-backend/post/internal/core/ports"
	"github.com/saaramahmoudi/twitter-backend/utils"
	"google.golang.org/api/iterator"
	"log"
)

type PostFirestore struct{

}

var client *firestore.Client
var ctx = context.Background()
var app *firebase.App
const CollectionAddress = "Posts"
const EventCollectionAddress = "Events"

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

func (repo PostFirestore)  GetSaveTransaction(id * string, operation func (* domain.Post) (* domain.Post, error)) (* domain.Post, error){
	ref := client.Collection(CollectionAddress).Doc(*id)
	post := &domain.Post{}
	err := client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		doc, err := tx.Get(ref) // tx.Get, NOT ref.Get!
		if err != nil {
			return err
		}
		err = doc.DataTo(post)
		if err != nil {
			return err
		}
		post, err = operation(post)
		if err != nil {
			return err
		}
		mapSave, err := utils.TurnStructToMap(post)
		if err != nil {
			return err
		}

		return tx.Set(ref, mapSave)
	})
	return post, err
}
func (repo PostFirestore) GetEventId (event * domain.PostEvent) (* string, error) {
	//TODO this way of getting the id can cause fracture in change, change it in future
	log.Println(*event.MadeByUserId, *event.PostId, event.EventType)
	iter := client.Collection(EventCollectionAddress).Where("madeByUserId", "==", *event.MadeByUserId).
												Where("postId", "==", *event.PostId).
												 Where("eventType", "==", event.EventType).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		event := domain.PostEvent{}

		err = doc.DataTo(&event)
		if err != nil {
			log.Println(err)
			continue
		}
		return event.Id, nil
	}
	return nil, errors.New(ports.NoEventFound)
}
func (repo PostFirestore) SaveOrDeleteEvent(event * domain.PostEvent) error {
	if *event.IsReversal {

		id, err := repo.GetEventId(event)
		if err != nil {
			return err
		}
		event.Id = id
		_, err = client.Collection(EventCollectionAddress).Doc(*event.Id).Delete(ctx)
		return err
	}else {
		id, err := repo.GetNewEventId()
		if err != nil{
			return err
		}
		event.Id = id
		mapSave, err := utils.TurnStructToMap(event)
		if err != nil{
			return err
		}
		_, err = client.Collection(EventCollectionAddress).Doc(*event.Id).Set(ctx, mapSave)
		return err
	}
}
func (repo PostFirestore) GetNewEventId() (*string, error) {
	ref := client.Collection(EventCollectionAddress).NewDoc()
	if ref.ID == ""{
		return nil, errors.New("Could not create a new post")
	}
	return &ref.ID, nil
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









