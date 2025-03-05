package models

import "gorm.io/gorm"

type Merchant struct {
	gorm.Model
	// ID           int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Nickname     string `gorm:"column:nickname;not null;comment:昵称" json:"nickname"` // 昵称
	Avatar       string `gorm:"column:avatar;comment:头像" json:"avatar"`              // 头像
	Introduction string `gorm:"column:introduction;comment:简介" json:"introduction"`  // 简介
	Password     string `gorm:"column:password;not null" json:"password"`
	Email        string `gorm:"column:email;not null" json:"email"`
	Phone        string `gorm:"column:phone" json:"phone"`
	CreatedBy    string `gorm:"column:created_by" json:"created_by"`
}

func (Merchant) TableName() string {
	return "merchant"
}