package handler

import (
	"context"
	"encoding/json"
	"github.com/saaramahmoudi/twitter-backend/post/internal/core/service"
	"github.com/saaramahmoudi/twitter-backend/post/pkg/errors"
	"github.com/saaramahmoudi/twitter-backend/tweet"
	"github.com/saaramahmoudi/twitter-backend/user"
	"log"
	"net/http"
)

// TODO do factory for this later
type HttpHandler struct {
	PostService service.PostService
	AuthService user.UserAuth
}
// TODO check if this can be factored out into only being into the users module
func (handler * HttpHandler) GetAuthContext(w http.ResponseWriter, req * http.Request) context.Context{
	header := req.Header.Get("Authorization")
	return context.WithValue(context.Background(), handler.AuthService.GetAuthHeaderKey(), header)
}
type GetPostByIdInput struct {
	Id string `json:"id"`
}
func (handler * HttpHandler) GetPostById(w http.ResponseWriter, req * http.Request){

	reqT := GetPostByIdInput{}
	err := json.NewDecoder(req.Body).Decode(&reqT)
	if err != nil {
		log.Fatal(err)
	}
	id := reqT.Id
	post, err := handler.PostService.Get(handler.GetAuthContext(w, req), &id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		json.NewEncoder(w).Encode(errors.UserErrors{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(post)
}

type CreatePostInput struct {
	Text * string `json:"text"`
	MediaType * tweet.MediaType `json:"mediaType"`
}

func (handler * HttpHandler) CreatePost(w http.ResponseWriter, req * http.Request){


	reqT := CreatePostInput{}
	err := json.NewDecoder(req.Body).Decode(&reqT)
	if err != nil {
		log.Fatal(err)
	}
	post, err := handler.PostService.Create(handler.GetAuthContext(w, req), reqT.Text, reqT.MediaType)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		json.NewEncoder(w).Encode(errors.UserErrors{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(post)
}



