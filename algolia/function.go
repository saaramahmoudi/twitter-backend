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

type DataStoreUser struct {
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

func getObject(fields interface{}, transferToData interface {}) (error) {

	finalMap := make(map[string]interface{})

	log.Println("Got this data : ", fields)
	//TODO find a better way to transform this object
	for key, element := range fields.(map[string]interface{}) {
		for key2, element2 := range element.(map[string]interface{}){
			if key2 == "arrayValue" {
			}else {
				finalMap[key] = element2
			}
		}
	}

	log.Println("Got this map : ", finalMap)

	bytes, err := json.Marshal(finalMap)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, transferToData)
	if err != nil {
		return err
	}

	return nil
}
var client * search.Client
func init(){
	client = search.NewClient(os.Getenv("API_ID"), os.Getenv("API_KEY"))
}

type FinalData struct {
	ObjectID string `json:"objectID"`
	Fields     interface{} `json:"fields"`
}

func getIdFromField(fields interface{}) string {
	map1 := fields.(map[string]interface{})["id"]
	res := map1.(map[string]interface{})["stringValue"]
	return res.(string)
}

func FireStoreTransfer(ctx context.Context, e FirestoreEvent, index * search.Index) error {

	var err error

	if e.Value.Fields == nil {
		log.Println("got this value : ", e.OldValue.Fields, e.OldValue.Name)
		log.Println("Id : ", getIdFromField(e.OldValue.Fields))

		d, err := index.DeleteObject(getIdFromField(e.OldValue.Fields))
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println(d)
		return nil
	}

	log.Println("got this value : ", e.Value.Fields, e.Value.Name)
	log.Println("Id : ", getIdFromField(e.Value.Fields))

	ret := FinalData{ObjectID: getIdFromField(e.Value.Fields), Fields: e.Value.Fields}
	log.Println(ret)
	res, err := index.SaveObject(ret)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(res)
	return nil
}
func FireStoreTransferUser(ctx context.Context, e FirestoreEvent) error {
	index := client.InitIndex("user")

	return FireStoreTransfer(ctx, e, index)
}

func FireStoreTransferTweet(ctx context.Context, e FirestoreEvent) error {
	index := client.InitIndex("tweet")

	return FireStoreTransfer(ctx, e, index)
}