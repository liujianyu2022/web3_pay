package models

import "gorm.io/gorm"

type SysWallet struct {
	gorm.Model
	// ID           int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Ac           string `gorm:"column:ac;not null;comment:钱包地址" json:"ac"`                         // 钱包地址
	PriKey       string `gorm:"column:pri_key;not null;comment:私钥" json:"pri_key"`                 // 私钥
	Mnemonic     string `gorm:"column:mnemonic;not null;comment:助记词" json:"mnemonic"`              // 助记词
	Path         string `gorm:"column:path;not null;comment:钱包路径" json:"path"`                     // 钱包路径
	Remark       string `gorm:"column:remark;comment:备注" json:"remark"`                            // 备注
	CurrentIndex int32  `gorm:"column:current_index;not null;comment:当前派生索引" json:"current_index"` // 当前派生索引
	Version      int32  `gorm:"column:version;not null" json:"version"`
}

func (SysWallet) TableName() string {
	return "sys_wallet"
}