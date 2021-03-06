package repositories

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	firebase "firebase.google.com/go"
	"github.com/saaramahmoudi/twitter-backend/user/internal/core/domain"
	"github.com/saaramahmoudi/twitter-backend/utils"
	"google.golang.org/api/iterator"
	"log"
)

type UserFirestore struct{

}

var client *firestore.Client
var ctx = context.Background()
var app *firebase.App
const CollectionAddress = "UserProfile"
// This is used because it seems that the gcp firestore library does not use json tags correctly for creating documents

const EmailNotExists = "user not found"
// TODO EmailNotExists creates an unseen dep between emailExists and get by email, remove in future iteration
func (repo UserFirestore) EmailExists(email * string) (*bool, error){
	_, err := repo.GetByEmail(email)
	found := true
	if err != nil && err.Error() == EmailNotExists {
		found = false
		err = nil
	} else if err != nil {
		found = false
	}

	return &found, err

}
func (repo UserFirestore) GetByEmail(email * string) (*domain.User, error){
	iter := client.Collection(CollectionAddress).Where("email", "==", *email).Documents(ctx)
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
		return &user, nil
	}
	return nil, errors.New(EmailNotExists)

}
func (repo UserFirestore) GetById(id * string) (*domain.User, error){
	doc, err := client.Collection(CollectionAddress).Doc(*id).Get(ctx)
	if err != nil {
		return nil, err
	}

	res := domain.User{}
	errDataTransfer := doc.DataTo(&res)

	if errDataTransfer != nil {
		return nil, errDataTransfer
	}
	// This is because we don't keep the email on the firestore to avoid repetition TODO check if it's a good design
	res.Email = id
	return &res, nil
}

func (repo UserFirestore) GetUserFromTag(tag * string) (* domain.User, error){
	iter := client.Collection(CollectionAddress).Where("tag", "==", tag).Documents(ctx)
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
		//This is because the tags are also unique to users, TODO remove this dep
		return &user, nil
	}
	return nil, errors.New("Couldn't find the user")
}
func (repo UserFirestore) UpdateUser(user * domain.User) (* domain.User, error) {
	mapUser, err := utils.TurnStructToMap(user)
	if err != nil{
		return nil, err
	}
	_, err = client.Collection(CollectionAddress).Doc(* user.Id).Set(ctx, mapUser)
	if err != nil{
		return nil, err
	}
	return user, nil
}
func checkTransactionAndGetData(user * domain.User, ref *firestore.DocumentRef, tx *firestore.Transaction) error{

	doc, err := tx.Get(ref) // tx.Get, NOT ref.Get!
	if err != nil {
		return err
	}
	err = doc.DataTo(user)
	return err
}
func saveTransaction(user * domain.User, ref *firestore.DocumentRef, tx *firestore.Transaction) error {
	mapSave, err := utils.TurnStructToMap(user)
	if err != nil {
		return err
	}
	return tx.Set(ref, mapSave)
}
func (repo UserFirestore) GetSaveTransactionTwoUsers(id1 * string, id2 * string, operation func (user1 * domain.User, user2 * domain.User) (* domain.User, * domain.User,  error)) (* domain.User, * domain.User, error){
	ref1 := client.Collection(CollectionAddress).Doc(*id1)
	ref2 := client.Collection(CollectionAddress).Doc(*id2)
	user1 := &domain.User{}
	user2 := &domain.User{}
	err := client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		err := checkTransactionAndGetData(user1, ref1, tx)
		if err != nil{
			return nil
		}
		err = checkTransactionAndGetData(user2, ref2, tx)
		if err != nil{
			return nil
		}
		user1, user2, err = operation(user1, user2)
		if err != nil {
			return err
		}
		err = saveTransaction(user2, ref2, tx)
		if err != nil {
			return err
		}

		return saveTransaction(user1, ref1, tx)
	})
	return user1, user2, err
}
//TODO check if we need to merge update and save
func (repo UserFirestore) Save(user * domain.User) (* domain.User, error){
	mapUser, err := utils.TurnStructToMap(user)
	if err != nil{
		return nil, err
	}
	_, err = client.Collection(CollectionAddress).Doc(* user.Id).Set(ctx, mapUser)
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









