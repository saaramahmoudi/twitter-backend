package handler

import (
	"context"
	"encoding/json"
	"github.com/saaramahmoudi/twitter-backend/user/internal/core/ports"
	"github.com/saaramahmoudi/twitter-backend/user/pkg/errors"
	"log"
	"net/http"
)

// TODO do factory for this later
type HttpHandler struct {
	UserService ports.UserService
	AuthService ports.HttpUserAuthenticator
}

func (handler * HttpHandler) GetAuthContext(w http.ResponseWriter, req * http.Request) context.Context{
	header := req.Header.Get("Authorization")
	return context.WithValue(context.Background(), handler.AuthService.GetAuthHeaderKey(), header)
}

type GetUserRequest struct {
	Email string  `json:"email"`
}

func (handler * HttpHandler) GetUser(w http.ResponseWriter, req * http.Request){

	reqT := GetUserRequest{}
	err := json.NewDecoder(req.Body).Decode(&reqT)
	if err != nil {
		log.Fatal(err)
	}
	email := reqT.Email
	user, err := handler.UserService.GetByEmail(handler.GetAuthContext(w, req), &email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		json.NewEncoder(w).Encode(errors.UserErrors{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(user)
}

type TagUpdateInput struct {
	Tag string `json:"tag"`
}

func (handler * HttpHandler) UpdateUserTag(w http.ResponseWriter, req * http.Request){
	in := TagUpdateInput{}
	err := json.NewDecoder(req.Body).Decode(&in)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.UserErrors{Message: err.Error()})
		return
	}
	user, err := handler.UserService.UpdateTag(handler.GetAuthContext(w, req), &in.Tag)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.UserErrors{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(user)
}


func (handler * HttpHandler) CreateUser(w http.ResponseWriter, req * http.Request){

	user, err := handler.UserService.Create(handler.GetAuthContext(w, req))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		json.NewEncoder(w).Encode(errors.UserErrors{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(user)
}

type DocCheckOutput struct{
	Exists bool `json:"exists"`
}
func (handler * HttpHandler) CheckDoc(w http.ResponseWriter, req * http.Request){

	exists, err := handler.UserService.EmailExists(handler.GetAuthContext(w, req))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		json.NewEncoder(w).Encode(errors.UserErrors{Message: err.Error()})
		return
	}
	res := DocCheckOutput{Exists: *exists}
	json.NewEncoder(w).Encode(&res)
}






