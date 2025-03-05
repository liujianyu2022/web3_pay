package config

import (
	"context"
	"fmt"
	"log"
	"time"
	"web3_pay_backend/models"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		WholeConfig.DatabaseConfig.User,
		WholeConfig.DatabaseConfig.Password,
		WholeConfig.DatabaseConfig.Host,
		WholeConfig.DatabaseConfig.Port,
		WholeConfig.DatabaseConfig.DatabaseName,
	)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移模型
	// GORM会根据提供的模型（User、Goods、Order）在数据库中创建相应的表（如果这些表尚不存在），或者更新这些表的结构以匹配模型定义（如添加新字段）
	DB.AutoMigrate(
		&models.User{},
		
		&models.Order{},

		&models.Merchant{},
		&models.MerchantAPI{},
		&models.MerchantMeta{},

		&models.Wallet{},
		&models.SysWallet{},

		&models.SysConfig{},
	)

	return DB
}

func NewRedis() *redis.Client {
	Redis := redis.NewClient(&redis.Options{
		Addr:     WholeConfig.RedisConfig.Addr,
		Password: WholeConfig.RedisConfig.Password,
		DB:       WholeConfig.RedisConfig.DB,			// 要连接的 Redis 数据库编号（Redis 默认有 16 个数据库，编号从 0 到 15）
	})

	// 创建一个带有超时时间的上下文 ctx，超时时间设置为 5 秒
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)		
	
	// 使用 defer cancel() 确保在函数返回前取消上下文，释放相关资源
	defer cancel()

	// 使用 Redis.Ping(ctx) 方法发送一个 PING 命令到 Redis 服务器，以测试连接是否成功
	// 通过调用 .Result() 方法获取 PING 命令的响应结果。如果连接成功，PING 命令通常会返回一个简单的 "PONG" 响应。
	_, err := Redis.Ping(ctx).Result()												

	if err != nil {
		panic(fmt.Sprintf("redis error: %s", err.Error()))
	}

	return Redis
}

func NewLog() *log.Logger {
	return log.Default()
}