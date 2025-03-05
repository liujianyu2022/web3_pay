//go:build wireinject
// +build wireinject

package wire

import (
	"web3_pay_backend/config"
	"web3_pay_backend/handlers"
	"web3_pay_backend/repository"
	"web3_pay_backend/router"
	"web3_pay_backend/services"
	"web3_pay_backend/tools"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var repositorySet = wire.NewSet(
	repository.NewRepository,
	repository.NewUserRepository,
)

var toolsSet = wire.NewSet(
	tools.NewSonyID,
	tools.NewJwt,
)

var configSet = wire.NewSet(
	config.NewDB,
	config.NewLog,
	config.NewRedis,
)

var serviceSet = wire.NewSet(
	services.NewService,
	services.NewUserService,
	
)

var handlerSet = wire.NewSet(
	handlers.NewUserHandler,
)

var routerSet = wire.NewSet(
	router.SetupRouter,
)

var SuperSet = wire.NewSet(
	configSet,
	toolsSet,
	routerSet,

	handlerSet,
	serviceSet,
	repositorySet,
)

func InitailizeApp() *gin.Engine {
	wire.Build(SuperSet)
	return &gin.Engine{}
}