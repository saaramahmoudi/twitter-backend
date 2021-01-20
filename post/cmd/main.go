package main

import (
	"github.com/gorilla/mux"
	"github.com/saaramahmoudi/twitter-backend/user"
	"net/http"
)

func main(){


	r := mux.NewRouter()
	r.HandleFunc("/update/{email}/{id}", user.UpdateUserIdFunction).Methods("PUT")
	r.HandleFunc("/{email}", user.GetUserFunction).Methods("GET")

	http.ListenAndServe(":8080", r)

}


