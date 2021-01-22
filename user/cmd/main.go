package main

import (
	"github.com/gorilla/mux"
	"github.com/saaramahmoudi/twitter-backend/user"
	"net/http"
)

func main(){


	r := mux.NewRouter()
	r.HandleFunc("/create", user.CreateUser)
	r.HandleFunc("/update", user.UpdateUserIdFunction)
	r.HandleFunc("/get", user.GetUserFunction)
	r.HandleFunc("/check", user.CheckDoc)

	http.ListenAndServe(":8080", r)

}


