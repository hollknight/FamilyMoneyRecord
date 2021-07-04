package account

import (
	"FamilyMoneyRecord/database/action/operation"
	"FamilyMoneyRecord/database/action/stock"
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
func AddAccount(db *gorm.DB, userID uint64) (uint64, error) {
	account := model.Account{
		UserID: userID,
	}

	err := db.Create(&account).Error

	return account.ID, err
}

// AddAccountByStruct 根据结构体添加证券账户
func AddAccountByStruct(db *gorm.DB, account model.Account) error {
	err := db.Create(&account).Error

	return err
}

// GetAllAccounts 获取股票账户列表
func GetAllAccounts(db *gorm.DB) ([]model.Account, error) {
	var accountsList []model.Account
	result := db.Find(&accountsList)

	return accountsList, result.Error
}

// GetAccountsByUserID 获取用户股票账户列表
func GetAccountsByUserID(db *gorm.DB, userID uint64) ([]model.Account, error) {
	var accountsList []model.Account
	result := db.Where("user_id = ?", userID).Find(&accountsList)

	return accountsList, result.Error
}

// GetAccountByID 根据id获取股票账户
func GetAccountByID(db *gorm.DB, id uint64) (model.Account, error) {
	account := new(model.Account)
	err := db.Where("id = ?", id).First(account).Error

	return *account, err
}

// UpdateAccountProfit 更新并获取股票账户盈亏金额
//func UpdateAccountProfit(db *gorm.DB, id uint64, profit float64) error {
//	account := new(model.Account)
//	res := db.Model(&account).Where("id = ?", id).Update("profit", profit)
//	if res.RowsAffected == 0 {
//		err := errors.New("该主键下股票账户已被删除")
//		return err
//	}
//
//	return res.Error
//}

// DeleteAccount 删除指定证券账户
func DeleteAccount(db *gorm.DB, id uint64) error {
	account := new(model.Account)
	err := db.Where("id = ?", id).Delete(&account).Error

	return err
}

// DeleteAccounts 删除列表中所有账单
func DeleteAccounts(db *gorm.DB, accountList []model.Account) error {
	for _, account := range accountList {
		operationList, err := operation.GetAllOperationsByAccountID(db, account.ID)
		if err != nil {
			return err
		}
		err = operation.DeleteOperations(db, operationList)
		if err != nil {
			return err
		}

		stockList, err := stock.GetStocksByAccountID(db, account.ID)
		if err != nil {
			return err
		}
		err = stock.DeleteStocks(db, stockList)
		if err != nil {
			return err
		}

		err = db.Delete(&account).Error
		if err != nil {
			return err
		}
	}
	return nil
}
