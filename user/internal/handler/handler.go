package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/saaramahmoudi/twitter-backend/user/internal/core/ports"
	"github.com/saaramahmoudi/twitter-backend/user/pkg/errors"
	"log"
	"net/http"
)

type HttpHandler struct {
	UserService ports.UserService
}

func (handler * HttpHandler) GetUser(w http.ResponseWriter, req * http.Request){
	vars := mux.Vars(req)
	email := vars["email"]
	user, err := handler.UserService.Get(email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		json.NewEncoder(w).Encode(errors.UserErrors{Message: "Could not find user email"})
	}
	json.NewEncoder(w).Encode(user)
}


func (handler * HttpHandler) UpdateUserId(w http.ResponseWriter, req * http.Request){
	vars := mux.Vars(req)
	email := vars["email"]
	id := vars["id"]
	user, err := handler.UserService.UpdateId(email, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		json.NewEncoder(w).Encode(errors.UserErrors{Message: "Could not find user email"})
	}
	json.NewEncoder(w).Encode(user)
}








