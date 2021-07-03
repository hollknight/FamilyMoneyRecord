package model

// Account 证券账户表
type Account struct {
	ID         uint64      `gorm:"primary_key;auto_increment" json:"id"`
	UserID     uint64      `json:"user_id"`
	Operations []Operation `json:"operations,omitempty"`
	Stocks     []Stock     `json:"stocks,omitempty"`
}
