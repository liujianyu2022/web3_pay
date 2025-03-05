package repository

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// 主要包括数据库（MySQL）和缓存（Redis）的初始化、事务管理以及日志集成

// 一个常量，用于标识事务的数据库连接
const ctxTxKey = "TxKey"

// 封装了数据库（gorm.DB）、缓存（redis.Client）和日志（log.Logger）的实例
type Repository struct {
	db      *gorm.DB
	redisDb *redis.Client
	logger  *log.Logger
}

// 用于创建Repository实例
func NewRepository(db *gorm.DB, redisDb *redis.Client, logger *log.Logger) *Repository {
	return &Repository{
		db:      db,
		redisDb: redisDb,
		logger:  logger,
	}
}

// 初始化MySQL数据库连接
// 返回当前上下文中的数据库连接，如果当前上下文中存在事务（ctxTxKey），则返回事务的数据库连接；否则返回普通的数据库连接
func (repository *Repository) DB(ctx *gin.Context) *gorm.DB {
	value := ctx.Value(ctxTxKey) // 从上下文中获取与ctxTxKey关联的值

	if value != nil { // 检查上下文中是否存在与ctxTxKey关联的值, 如果存在，说明当前处于事务中
		tx, ok := value.(*gorm.DB) // 将上下文中存储的值转换为*gorm.DB类型, 如果转换成功，说明上下文中存储的是一个事务的数据库连接
		if ok {
			return tx
		}
	}

	return repository.db.WithContext(ctx) // 如果上下文中不存在事务，返回普通的数据库连接
}



