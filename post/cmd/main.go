package main

import (
	"github.com/gorilla/mux"
	"github.com/saaramahmoudi/twitter-backend/post"
	"net/http"
)

func main(){


	r := mux.NewRouter()
	r.HandleFunc("/create", post.CreatePostFunction)
	r.HandleFunc("/like", post.ToggleLikePostFunction)
	r.HandleFunc("/retweet", post.ToggleRetweetPostFunction)
	r.HandleFunc("/get", post.GetPostByIdFunction)

	http.ListenAndServe(":8080", r)

}


