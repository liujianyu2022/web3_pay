package models

import "gorm.io/gorm"

type MerchantMeta struct {
	gorm.Model
	// ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	MID       int64          `gorm:"column:m_id;not null" json:"m_id"`
	Ac        string         `gorm:"column:ac;comment:钱包地址" json:"ac"` // 钱包地址
}

func (MerchantMeta) TableName() string {
	return "merchant_meta"
}