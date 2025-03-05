package models

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// ID           int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserType     string         `gorm:"column:user_type;not null" json:"user_type"`
	UserID       string         `gorm:"column:user_id;not null" json:"user_id"`
	Nickname     string         `gorm:"column:nickname;not null;comment:昵称" json:"nickname"` // 昵称
	Avatar       string         `gorm:"column:avatar;not null;comment:头像" json:"avatar"`     // 头像
	Introduction string         `gorm:"column:introduction;comment:简介" json:"introduction"`  // 简介
	Password     string         `gorm:"column:password;not null" json:"password"`
	Email        string         `gorm:"column:email; not null" json:"email"`
	Phone        string         `gorm:"column:phone" json:"phone"`
	LoginAt      time.Time      `gorm:"column:login_at; type:timestamp; default:CURRENT_TIMESTAMP" json:"login_at"`
}


func (user *User) TableName() string {
    return "user" 									// 指定表名为 "user"
}