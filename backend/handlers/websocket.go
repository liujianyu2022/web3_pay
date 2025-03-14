package handlers

import "github.com/gin-gonic/gin"

type WebsocketHandler struct {
}

func NewWebsocketHandler() WebsocketHandler {
	return WebsocketHandler{}
}

func (websocketHandler *WebsocketHandler) RunWebsocket(ctx *gin.Context) {

}
