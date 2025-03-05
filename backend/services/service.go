package services

import (
	"log"
	"web3_pay_backend/tools"
	// "github.com/golang-jwt/jwt/v5"
)

type Service struct {
	logger *log.Logger
	jwt    *tools.JWT
	sonyID *tools.SonyID
}

func NewService(logger *log.Logger, jwt *tools.JWT, sonyID *tools.SonyID) *Service {
	return &Service{
		logger: logger,
		jwt: jwt,
		sonyID: sonyID,
	}
}
