package model

// Stock 股票持仓表
type Stock struct {
	ID          uint64      `gorm:"primary_key;auto_increment" json:"id"`
	AccountID   uint64      `json:"account_id"`
	Code        string      `gorm:"type:varchar(10);unique;not null" json:"code"`
	PositionNum int         `gorm:"type:int" json:"position_num"`
	Profit      float64     `gorm:"type:double" json:"profit"`
	Operations  []Operation `json:"operations,omitempty"`
}
