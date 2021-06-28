package account

import (
	"FamilyMoneyRecord/database/model"
	"errors"
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
func UpdateAccountProfit(db *gorm.DB, id uint64, profit int) error {
	account := new(model.Account)
	res := db.Model(&account).Where("id = ?", id).Update("profit", profit)
	if res.RowsAffected == 0 {
		err := errors.New("该主键下股票账户已被删除")
		return err
	}

	return res.Error
}

// DeleteAccount 删除指定证券账户
func DeleteAccount(db *gorm.DB, id uint64) error {
	account := new(model.Account)
	err := db.Where("id = ?", id).First(account).Error

	return err
}

// DeleteAccounts 删除列表中所有账单
func DeleteAccounts(db *gorm.DB, accountList []model.Account) error {
	for _, account := range accountList {
		err := db.Delete(&account).Error
		if err != nil {
			return err
		}
	}
	return nil
}
