package stock

import (
	"FamilyMoneyRecord/database/model"
	"gorm.io/gorm"
)

//// Stock 股票持仓表
//type Stock struct {
//	ID          uint64 `gorm:"primary_key;auto_increment"`
//	AccountID   uint64
//	Code        string `gorm:"type:varchar(10);unique;not null"`
//	PositionNum int    `gorm:"type:int"`
//	Profit      int    `gorm:"type:int"`
//}

// AddStock 添加股票
func AddStock(db *gorm.DB, accountID uint64, code string, positionNum, profit int) (model.Stock, error) {
	stock := model.Stock{
		AccountID:   accountID,
		Code:        code,
		PositionNum: positionNum,
		Profit:      profit,
	}

	err := db.Create(&stock).Error

	return stock, err
}

// GetStock 获取股票信息
func GetStock(db *gorm.DB, accountID uint64, code string) (model.Stock, error) {
	stock := new(model.Stock)
	err := db.Where("account_id = ? AND code = ?", accountID, code).First(stock).Error

	return *stock, err
}

// GetStocksByAccountID 根据股票账户id获取所有股票
func GetStocksByAccountID(db *gorm.DB, accountID uint64) ([]model.Stock, error) {
	var stockList []model.Stock
	err := db.Where("account_id = ?", accountID).Find(stockList).Error

	return stockList, err
}

// UpdateStock 更新股票股数与盈利金额
func UpdateStock(db *gorm.DB, stock model.Stock, positionNum, profit int) error {
	err := db.Model(&stock).Updates(
		map[string]interface{}{
			"position_sum": positionNum,
			"profit":       profit,
		}).Error

	return err
}
