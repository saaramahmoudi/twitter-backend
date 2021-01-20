package main

import (
	"fmt"
	"github.com/saaramahmoudi/twitter-backend/tweet/pkg/handler"
)

func main(){


	fmt.Print(handler.TweetApi)
	//r := mux.NewRouter()
	//r.HandleFunc("/update/{email}/{id}", user.UpdateUserIdFunction).Methods("PUT")
	//r.HandleFunc("/{email}", user.GetUserFunction).Methods("GET")
	//
	//http.ListenAndServe(":8080", r)

}


