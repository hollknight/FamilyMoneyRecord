package operation

import (
	"FamilyMoneyRecord/database/model"
	"gorm.io/gorm"
	"time"
)

//// Operation 股票操作表
//type Operation struct {
//	ID         uint64 `gorm:"primary_key;auto_increment"`
//	AccountID  uint64
//	Code       string    `gorm:"type:varchar(10);unique;not null"`
//	SharePrice int       `gorm:"type:int"`
//	BuyNum     int       `gorm:"type:int;default:0"`
//	SaleNum    int       `gorm:"type:int;default:0"`
//	Time       time.Time `gorm:"type:datetime;default:'1000-01-01 00:00:00'"`
//}

// AddOperation 添加股票操作记录
func AddOperation(db *gorm.DB, accountID, stockID uint64, sharePrice float64, buyNum, SaleNum int) (uint64, error) {
	now := time.Now().Add(time.Hour * 8)

	operation := model.Operation{
		AccountID:  accountID,
		StockID:    stockID,
		SharePrice: sharePrice,
		BuyNum:     buyNum,
		SaleNum:    SaleNum,
		Time:       now,
	}

	err := db.Create(&operation).Error

	return operation.ID, err
}

// AddOperationByStruct 根据结构体添加股票操作记录
func AddOperationByStruct(db *gorm.DB, operation model.Operation) error {
	err := db.Create(&operation).Error

	return err
}

// GetAllOperations 获取所有交易记录
func GetAllOperations(db *gorm.DB) ([]model.Operation, error) {
	var operationList []model.Operation
	result := db.Find(&operationList)

	return operationList, result.Error
}

// GetAllOperationsByAccountID 获取股票账户下所有交易记录
func GetAllOperationsByAccountID(db *gorm.DB, accountID uint64) ([]model.Operation, error) {
	var operationList []model.Operation
	result := db.Where("account_id = ?", accountID).Find(&operationList)

	return operationList, result.Error
}

// GetOperationByID 根据主键获取股票交易记录
func GetOperationByID(db *gorm.DB, id uint64) (model.Operation, error) {
	operation := new(model.Operation)
	err := db.Where("id = ?", id).First(operation).Error

	return *operation, err
}

// DeleteOperationByID 删除交易记录
func DeleteOperationByID(db *gorm.DB, id uint64) (int, int, float64, error) {
	operation := new(model.Operation)
	err := db.Where("id = ?", id).First(operation).Error
	if err != nil {
		return 0, 0, 0, err
	}
	err = db.Delete(&operation).Error
	return operation.BuyNum, operation.SaleNum, operation.SharePrice, err
}

// UpdateOperationByID 修改交易记录
func UpdateOperationByID(db *gorm.DB, id uint64, buyNum, saleNum int, sharePrice float64) error {
	operation := new(model.Operation)
	_, err := GetOperationByID(db, id)
	if err != nil {
		return err
	}

	err = db.Model(&operation).Updates(
		map[string]interface{}{
			"buy_num":     buyNum,
			"sale_num":    saleNum,
			"share_price": sharePrice,
		}).Error

	return err
}
