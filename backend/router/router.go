package router

import (
	"web3_pay_backend/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	userHandler *handlers.UserHandler,
	websocketHandler *handlers.WebsocketHandler,
) *gin.Engine {
	ginServer := gin.Default()

	api := ginServer.Group("/api")
	{
		api.POST("/register", userHandler.Register)
		api.POST("/login", userHandler.Login)
	}

	websocket := ginServer.Group("/websocket")
	{
		websocket.GET("/", websocketHandler.RunWebsocket)
	}

	return ginServer
}
