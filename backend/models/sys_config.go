package models

import "gorm.io/gorm"

type SysConfig struct {
	gorm.Model
	// ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Domain string `gorm:"column:domain;not null;comment:项目域名" json:"domain"` // 项目域名
}

func (SysConfig) TableName() string {
	return "sys_config"
}