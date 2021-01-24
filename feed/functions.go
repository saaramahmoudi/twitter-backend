package feed

import (
	"github.com/saaramahmoudi/twitter-backend/user"
	"github.com/saaramahmoudi/twitter-backend/feed/internal/handler"
)

var messageHandler = handler.MessageHandler{user.Api}
var Handle = messageHandler.Handle






