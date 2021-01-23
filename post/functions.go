package post

import (
	"github.com/saaramahmoudi/twitter-backend/post/internal/core/service"
	"github.com/saaramahmoudi/twitter-backend/post/internal/handler"
	"github.com/saaramahmoudi/twitter-backend/post/internal/repositories"
	"github.com/saaramahmoudi/twitter-backend/user"
	"net/http"
)
// TODO factor this out to the utils package
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

var authHandler = user.ApiAuth
var postService = service.PostService{Repo: repositories.PostFirestore{}}
var httpHandler = handler.HttpHandler{PostService: postService, AuthService: authHandler}
var GetPostByIdFunction = CORSCheck(httpHandler.GetPostById)
var CreatePostFunction = CORSCheck(httpHandler.CreatePost)





//gcloud functions deploy TestSara --entry-point GetUserFunction --runtime go113 --max-instances 2 --trigger-http --allow-unauthenticated