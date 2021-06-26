package model

// Stock 股票持仓表
type Stock struct {
	ID          uint64 `gorm:"primary_key;auto_increment"`
	AccountID   uint64
	Code        string `gorm:"type:varchar(10);unique;not null"`
	PositionNum int    `gorm:"type:int"`
	Profit      int    `gorm:"type:int"`
	Operations  []Operation
}
