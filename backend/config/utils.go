package config

import (
	"encoding/json"
	"os"
)

type DatabaseConfigType struct {
	User         string `json:"user"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	DatabaseName string `json:"database_name"`
}

type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type JwtConfig struct {
	SecurityKey string `json:"security_key"`
}

type WholeConfigType struct {
	DatabaseConfig DatabaseConfigType `json:"database"`
	RedisConfig    RedisConfig        `json:"redis"`
	JwtConfig      JwtConfig          `json:"jwt"`
}

var WholeConfig WholeConfigType

func LoadConfigFile(filename string) error {
	data, err := os.ReadFile(filename)

	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &WholeConfig) // 解析JSON数据到配置结构体

	if err != nil {
		return err
	}

	return nil
}
