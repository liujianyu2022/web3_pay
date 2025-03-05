package tools

import (
	"errors"
	"strings"
	"time"
	"web3_pay_backend/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type JWT struct {
	key   []byte
	redis *redis.Client
}

type CustomClaims struct {
	UserId string
	UserType string
	ClientIP string
	jwt.RegisteredClaims
}

func NewJwt(redis *redis.Client) *JWT {
	return &JWT{
		key: []byte(config.WholeConfig.JwtConfig.SecurityKey),
		redis: redis,
	}
}


func (j *JWT) GenToken(userId, userType, clientIP string, expiresAt time.Time) (string, error) {
	claims := CustomClaims{
		UserId:   userId,
		UserType: userType,
		ClientIP: clientIP,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token.Header["cty"] = "web3_pay_authority"

	return token.SignedString(j.key)
}

func (j *JWT) ParseToken(ctx *gin.Context, tokenString string) (*CustomClaims, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	
	if tokenString == "" {
		return nil, errors.New("token is empty")
	}

	if j.IsInBlacklist(ctx, tokenString) {
		return nil, errors.New("token revoked")
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		return j.key, nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func (j *JWT) AddToBlacklist(ctx *gin.Context, tokenString string, expiresAt *jwt.NumericDate) error {
	if expiresAt == nil {
		return errors.New("invalid token expiration")
	}

	ttl := time.Until(expiresAt.Time)
	if ttl <= 0 {
		return nil
	}

	return j.redis.SetNX(ctx, "jwt:blacklist:"+tokenString, "1", ttl).Err()
}

func (j *JWT) IsInBlacklist(ctx *gin.Context, tokenString string) bool {
	exists, _ := j.redis.Exists(ctx, "jwt:blacklist:"+tokenString).Result()
	return exists > 0
}