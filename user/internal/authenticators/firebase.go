package authenticators

import (
	"context"
	"errors"
	firebase "firebase.google.com/go"
	"log"
	"strings"
)

type FirebaseAuthenticator struct {

}


const RequestKey = "Request"
var ctx = context.Background()
var app *firebase.App
const Bearer = "Bearer "

func (fa FirebaseAuthenticator ) GetAuthHeaderKey() string {
	return "Auth"
}

func (fa FirebaseAuthenticator ) getClaims(ctxReq context.Context) (map[string]interface{}, error){
	client, err := app.Auth(ctx)

	if err != nil {
		return nil, err
	}

	authorHeader := ctxReq.Value(fa.GetAuthHeaderKey()).(string)
	if ! strings.HasPrefix(authorHeader, Bearer) {
		return nil, errors.New("Bad Auth header")
	}

	idToken := authorHeader[len(Bearer):]

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, err
	}
	return token.Claims, nil
}


func (fa FirebaseAuthenticator ) GetUid(ctxReq context.Context) (* string, error) {

	claims, err :=  fa.getClaims(ctxReq)
	if err != nil {
		return nil, err
	}

	uid, exists := claims["uid"].(string)
	if !exists {
		return nil, errors.New("Not a valid claim for uid")
	}
	return &uid, nil
}

func (fa FirebaseAuthenticator ) GetEmail(ctxReq context.Context) (* string, error){
	claims, err :=  fa.getClaims(ctxReq)
	if err != nil {
		return nil, err
	}

	email, exists := claims["email"].(string)
	if !exists {
		return nil, errors.New("Not a valid claim for email")
	}

	return &email, nil
}
func (fa FirebaseAuthenticator ) IsLoggedIn(ctxReq context.Context) bool {
	_, err := fa.getClaims(ctxReq)
	if err != nil {
		return false
	}
	return true
}


func init(){
	var err error

	app, err = firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}
}






















