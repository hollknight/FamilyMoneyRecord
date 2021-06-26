package model

import "time"

// Operation 股票操作表
type Operation struct {
	ID         uint64 `gorm:"primary_key;auto_increment"`
	AccountID  uint64
	StockID    uint64
	SharePrice int       `gorm:"type:int"`
	BuyNum     int       `gorm:"type:int;default:0"`
	SaleNum    int       `gorm:"type:int;default:0"`
	Time       time.Time `gorm:"type:datetime;default:'1000-01-01 00:00:00'"`
}
