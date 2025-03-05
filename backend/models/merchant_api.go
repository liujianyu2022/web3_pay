package models

import "gorm.io/gorm"

type MerchantAPI struct {
	gorm.Model
	// ID          int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	MID         int64          `gorm:"column:m_id;not null" json:"m_id"`
	Apikey      string         `gorm:"column:apikey;not null;comment:Token" json:"apikey"`             // Token
	CallbackURL string         `gorm:"column:callback_url;not null;comment:回调url" json:"callback_url"` // 回调url
	SecretKey   string         `gorm:"column:secret_key;not null;comment:回调加密参数" json:"secret_key"`    // 回调加密参数
	Remark      string         `gorm:"column:remark;comment:备注" json:"remark"`                         // 备注
}

func (MerchantAPI) TableName() string {
	return "merchant_api"
}