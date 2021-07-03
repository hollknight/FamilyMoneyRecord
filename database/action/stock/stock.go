package stock

import (
	"FamilyMoneyRecord/database/model"
	"gorm.io/gorm"
)

// AddStock 添加股票
func AddStock(db *gorm.DB, accountID uint64, code string, positionNum int, profit float64) (model.Stock, error) {
	stock := model.Stock{
		AccountID:   accountID,
		Code:        code,
		PositionNum: positionNum,
		Profit:      profit,
	}

	err := db.Create(&stock).Error

	return stock, err
}

// AddStockByStruct 根据结构体添加股票
func AddStockByStruct(db *gorm.DB, stock model.Stock) error {
	err := db.Create(&stock).Error

	return err
}

// GetAllStocks 获取股票信息
func GetAllStocks(db *gorm.DB) ([]model.Stock, error) {
	var stockList []model.Stock
	result := db.Find(&stockList)

	return stockList, result.Error
}

// GetStock 获取股票信息
func GetStock(db *gorm.DB, accountID uint64, code string) (model.Stock, error) {
	stock := new(model.Stock)
	err := db.Where("account_id = ? AND code = ?", accountID, code).First(stock).Error

	return *stock, err
}

// GetStockByID 根据主键获取股票信息
func GetStockByID(db *gorm.DB, id uint64) (model.Stock, error) {
	stock := new(model.Stock)
	err := db.Where("id = ?", id).First(stock).Error

	return *stock, err
}

// GetStocksByAccountID 根据股票账户id获取所有股票
func GetStocksByAccountID(db *gorm.DB, accountID uint64) ([]model.Stock, error) {
	var stockList []model.Stock
	err := db.Where("account_id = ?", accountID).Find(&stockList).Error

	return stockList, err
}

// UpdateStock 更新股票股数与盈利金额
func UpdateStock(db *gorm.DB, stock model.Stock, positionNum int, profit float64) error {
	err := db.Model(&stock).Updates(
		map[string]interface{}{
			"position_num": positionNum,
			"profit":       profit,
		}).Error

	return err
}

//
