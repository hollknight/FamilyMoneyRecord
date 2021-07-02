package model

import "time"

// Operation 股票操作表
type Operation struct {
	ID         uint64    `gorm:"primary_key;auto_increment" json:"id"`
	AccountID  uint64    `json:"accountID"`
	StockID    uint64    `json:"stock_id"`
	SharePrice float64   `gorm:"type:double" json:"share_price"`
	BuyNum     int       `gorm:"type:double;default:0" json:"buy_num"`
	SaleNum    int       `gorm:"type:double;default:0" json:"sale_num"`
	Time       time.Time `gorm:"type:datetime;default:'1000-01-01 00:00:00'" json:"time"`
}
