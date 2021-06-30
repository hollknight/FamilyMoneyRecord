package model

import "time"

// Operation 股票操作表
type Operation struct {
	ID         uint64 `gorm:"primary_key;auto_increment"`
	AccountID  uint64
	StockID    uint64
	SharePrice float64   `gorm:"type:double"`
	BuyNum     float64   `gorm:"type:double;default:0"`
	SaleNum    float64   `gorm:"type:double;default:0"`
	Time       time.Time `gorm:"type:datetime;default:'1000-01-01 00:00:00'"`
}
