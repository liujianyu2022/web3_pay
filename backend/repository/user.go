package repository

import (
	"errors"
	"web3_pay_backend/api"
	"web3_pay_backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Create(ctx *gin.Context, user *models.User) error				// 创建新用户
	Update(ctx *gin.Context, user *models.User) error				// 更新用户信息
	GetByID(ctx *gin.Context, id string) (*models.User, error)		// 根据用户ID查询单个用户信息
	GetByEmail(ctx *gin.Context, email string) (*models.User, error)
}

type UserRepository struct {
	repository *Repository
}

func NewUserRepository(repository *Repository) *UserRepository {
	return &UserRepository{
		repository: repository,
	}
}

func (userRepository *UserRepository) Create(ctx *gin.Context, user *models.User) error {
	err := userRepository.repository.DB(ctx).Create(user).Error

	if err != nil {
		return err
	}

	return nil
}

func (userRepository *UserRepository) Update(ctx *gin.Context, user *models.User) error {
	err := userRepository.repository.DB(ctx).Save(user).Error

	if err != nil {
		return err
	}

	return nil
}

func (userRepository *UserRepository) GetByID(ctx *gin.Context, id string) (*models.User, error) {
	var user *models.User
	err := userRepository.repository.DB(ctx).Where("user_id = ?", id).First(user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, api.ErrNotFound
		}
		return nil, err
	}

	return user, nil
}

func (userRepository *UserRepository) GetByEmail(ctx *gin.Context, email string) (*models.User, error) {
	var user models.User

	err := userRepository.repository.DB(ctx).Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, api.ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}
