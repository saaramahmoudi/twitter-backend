package user

import (
	"github.com/saaramahmoudi/twitter-backend/user/internal/authenticators"
	"github.com/saaramahmoudi/twitter-backend/user/internal/core/service"
	"github.com/saaramahmoudi/twitter-backend/user/internal/handler"
	"github.com/saaramahmoudi/twitter-backend/user/internal/repositories"
	"net/http"
)

func CORSCheck(handler func (w http.ResponseWriter, req *http.Request)) func (w http.ResponseWriter, req *http.Request){
	res := func (w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Max-Age", "3600")
		if req.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		handler(w, req)
	}
	return res
}
var authHandler = authenticators.FirebaseAuthenticator{}
var userService = service.UserService{Repo: repositories.UserFirestore{}, Auth: authHandler}
var httpHandler = handler.HttpHandler{UserService: userService, AuthService: authHandler}
var GetUserFunction = CORSCheck(httpHandler.GetUser)
var UpdateUserIdFunction = CORSCheck(httpHandler.UpdateUserId)
var CreateUser = CORSCheck(httpHandler.CreateUser)
var CheckDoc = CORSCheck(httpHandler.CheckDoc)





//gcloud functions deploy TestSara --entry-point GetUserFunction --runtime go113 --max-instances 2 --trigger-http --allow-unauthenticated