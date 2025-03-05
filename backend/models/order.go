package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	// ID           int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	MID          int64  `gorm:"column:m_id;not null" json:"m_id"`
	No           string `gorm:"column:no;not null;comment:订单no" json:"no"`                // 订单no
	CNo          string `gorm:"column:c_no;not null;comment:用户no" json:"c_no"`            // 用户no
	Amount       string `gorm:"column:amount;not null;comment:价格" json:"amount"`          // 价格
	Coin         string `gorm:"column:coin;not null;comment:数字币类型" json:"coin"`           // 数字币类型
	Chain        string `gorm:"column:chain;not null;comment:链" json:"chain"`             // 链
	NotifyURL    string `gorm:"column:notify_url;comment:回调url" json:"notify_url"`        // 回调url
	ReturnURL    string `gorm:"column:return_url;comment:前端重定向url" json:"return_url"`     // 前端重定向url
	PayURL       string `gorm:"column:pay_url;comment:支付地址" json:"pay_url"`               // 支付地址
	TxHash       string `gorm:"column:tx_hash;comment:交易hash" json:"tx_hash"`             // 交易hash
	Ac           string `gorm:"column:ac;comment:钱包地址" json:"ac"`                         // 钱包地址
	APIKey       string `gorm:"column:api_key;comment:商户apikey" json:"api_key"`           // 商户apikey
	TimeOut      int32  `gorm:"column:time_out;comment:超时时间" json:"time_out"`             // 超时时间
	Status       string `gorm:"column:status;comment:订单状态" json:"status"`                 // 订单状态
	NotifyStatus string `gorm:"column:notify_status;comment:订单回调结果" json:"notify_status"` // 订单回调结果
	Remark       string `gorm:"column:remark;comment:备注" json:"remark"`                   // 备注
}

func (Order) TableName() string {
	return "order"
}