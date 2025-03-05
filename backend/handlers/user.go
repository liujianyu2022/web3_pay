package handlers

import (
	"log"
	"net/http"
	"web3_pay_backend/api"
	"web3_pay_backend/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (userhandler *UserHandler) Register(ctx *gin.Context) {
	var request api.RegisterRequest

	err := ctx.ShouldBindJSON(&request)

	log.Println("request = ", request)

	if err != nil {
		api.HandleError(ctx, http.StatusBadRequest, api.ErrBadRequest, nil)
		return
	}

	err = userhandler.userService.Register(ctx, &request)

	if err != nil {
		api.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	api.HandleSuccess(ctx, "register success!")
}

func (userhandler *UserHandler) Login(ctx *gin.Context) {
	var request api.LoginRequest

	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		api.HandleError(ctx, http.StatusBadRequest, api.ErrBadRequest, nil)
		return
	}

	// 获取客户端IP
	clientIP := api.GetClientIP(ctx.Request)

	token, err := userhandler.userService.Login(ctx, &request, clientIP, "admin")

	if err != nil {
		api.HandleError(ctx, http.StatusInternalServerError, api.ErrUnauthorizedAP, nil)
		return
	}

	api.HandleSuccess(ctx, api.LoginResponseData{
		AccessToken: token,
	})
}
