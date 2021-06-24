package account

import (
	operation2 "FamilyMoneyRecord/database/models/operation"
	stock2 "FamilyMoneyRecord/database/models/stock"
)

// Account 证券账户表
type Account struct {
	ID         uint64 `gorm:"primary_key;auto_increment"`
	UserID     uint64
	Profit     int `gorm:"type:int;default:0"`
	Operations []operation2.Operation
	Stocks     []stock2.Stock
}
