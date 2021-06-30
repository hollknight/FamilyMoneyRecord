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

// GetAllOperations 获取股票账户下所有交易记录
func GetAllOperations(db *gorm.DB, accountID uint64) ([]model.Operation, error) {
	var operationList []model.Operation
	result := db.Where("account_id = ?", accountID).Find(&operationList)

	return operationList, result.Error
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
