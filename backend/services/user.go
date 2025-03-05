package services

import (
	"time"
	"web3_pay_backend/api"
	"web3_pay_backend/models"
	"web3_pay_backend/repository"

	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	*Service
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository, service *Service) *UserService {
	return &UserService{
		userRepository: userRepository,
		Service: service,
	}
}

func (userService *UserService) Register(ctx *gin.Context, request *api.RegisterRequest) error {

	email, nickname, password, rePassword := request.Email, request.Nickname, request.Password, request.RePassword

	user, err := userService.userRepository.GetByEmail(ctx, email)

	if user != nil && err == nil {
		return api.ErrEmailAlreadyUse
	}

	if password != rePassword {
		return api.ErrRePasswd
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	userID, err := userService.sonyID.GenerateString()
	if err != nil {
		return err
	}

	user = &models.User{
		UserID:   userID,
		Nickname: nickname,
		Email:    email,
		Password: string(hashedPassword),
	}

	err = userService.userRepository.Create(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func (userService *UserService) Login(ctx *gin.Context, request *api.LoginRequest, clientIP string, userType string) (string, error) {
	var email string = request.Email
	var password string = request.Password

	user, err := userService.userRepository.GetByEmail(ctx, email)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token, err := userService.jwt.GenToken(user.UserID, user.UserType, clientIP, time.Now().Add(time.Hour))
	if err != nil {
		return "", err
	}

	return token, nil
}
