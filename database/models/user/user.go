package user

import (
	account2 "FamilyMoneyRecord/database/models/account"
	bill2 "FamilyMoneyRecord/database/models/bill"
)

// User 用户表
type User struct {
	ID                 uint64 `gorm:"primary_key;auto_increment"`
	Username           string `gorm:"type:varchar(100);not null;unique"`
	Password           string `gorm:"type:char(32);not null"`
	Name               string `gorm:"type:varchar(20);default:''"`
	ReceiptSum         int    `gorm:"type:int;default:0"`
	DisbursementSum    int    `gorm:"type:int;default:0"`
	AdvanceConsumption int    `gorm:"type:int;default:0"`
	Bills              []bill2.Bill
	Accounts           []account2.Account
}
