package router

import (
	"web3_pay_backend/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	userHandler *handlers.UserHandler,
) *gin.Engine {
	ginServer := gin.Default()

	api := ginServer.Group("/api")
	{
		api.POST("/register", userHandler.Register)
		api.POST("/login", userHandler.Login)
	}

	return ginServer
}
