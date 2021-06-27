package account

import (
	"FamilyMoneyRecord/database/model"
	"gorm.io/gorm"
)

//// Account 证券账户表
//type Account struct {
//	ID         uint64 `gorm:"primary_key;auto_increment"`
//	UserID     uint64
//	Profit     int `gorm:"type:int;default:0"`
//	Operations []operation2.Operation
//	Stocks     []stock2.Stock
//}

// AddAccount 添加证券账户
func AddAccount(db *gorm.DB, userID uint64) error {
	account := model.Account{
		UserID: userID,
		Profit: 0,
	}

	err := db.Create(&account).Error

	return err
}

// GetAccountsByUserID 获取用户股票账户列表
func GetAccountsByUserID(db *gorm.DB, userID uint64) ([]model.Account, error) {
	var accountsList []model.Account
	result := db.Where("user_id = ?", userID).Find(&accountsList)

	return accountsList, result.Error
}

// UpdateAccountProfit 更新并获取股票账户盈亏金额
func UpdateAccountProfit(db *gorm.DB, id uint64, profit int) (model.Account, error) {
	account := new(model.Account)
	err := db.Where("id = ?", id).First(account).Error
	if err != nil {
		return *account, err
	}
	err = db.Model(&account).Where("id = ?", id).Update("profit", profit).Error

	return *account, err
}

// DeleteAccount 删除指定证券账户
func DeleteAccount(db *gorm.DB, id uint64) error {
	account := new(model.Account)
	err := db.Where("id = ?", id).First(account).Error

	return err
}
