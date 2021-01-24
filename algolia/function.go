package algolia

import (
	"context"
	"encoding/json"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/saaramahmoudi/twitter-backend/user"
	"log"
	"os"
	"time"
)

type ResClass struct {
	ObjectID string `json:"objectID"`
	user.User
}


// FirestoreEvent is the payload of a Firestore event.
type FirestoreEvent struct {
	OldValue   FirestoreValue `json:"oldValue"`
	Value      FirestoreValue `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// FirestoreValue holds Firestore fields.
type FirestoreValue struct {
	CreateTime time.Time `json:"createTime"`
	// Fields is the data for this value. The type depends on the format of your
	// database. Log the interface{} value and inspect the result to see a JSON
	// representation of your database fields.
	Fields     interface{} `json:"fields"`
	Name       string      `json:"name"`
	UpdateTime time.Time   `json:"updateTime"`
}

// HelloFirestore is triggered by a change to a Firestore document.
func FireStoreTransfer(ctx context.Context, e FirestoreEvent) error {

	client := search.NewClient(os.Getenv("API_ID"), os.Getenv("API_KEY"))
	index := client.InitIndex("user")

	finalMap := make(map[string]interface{})
	for key, element := range e.Value.Fields.(map[string]interface{}) {
		for _, element2 := range element.(map[string]interface{}){
			finalMap[key] = element2
		}
	}
	bytes, err := json.Marshal(finalMap)
	if err != nil {
		log.Println(err)
		return err
	}

	userInstance := &user.User{}
	err = json.Unmarshal(bytes, userInstance)
	if err != nil {
		log.Println(err)
		return err
	}
	ret := ResClass{ObjectID: *userInstance.Id, User: *userInstance}
	_, err = index.SaveObject(ret)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}