package bill

import (
	"FamilyMoneyRecord/database/model"
	"gorm.io/gorm"
	"time"
)

//// Bill 账单表
//type Bill struct {
//	ID           uint64 `gorm:"primary_key;auto_increment"`
//	UserID       uint64
//	Receipt      int       `gorm:"type:int;default:0"`
//	Disbursement int       `gorm:"type:int;default:0"`
//	Type         string    `gorm:"type:varchar(25);default:''"`
//	Time         time.Time `gorm:"type:datetime;default:'1000-01-01 00:00:00'"`
//}

// AddBill 添加账单
func AddBill(db *gorm.DB, user model.User, receipt, disbursement float64, moneyType, note string) (uint64, error) {
	now := time.Now().Add(time.Hour * 8)

	bill := model.Bill{
		UserID:       user.ID,
		Receipt:      receipt,
		Disbursement: disbursement,
		Type:         moneyType,
		Note:         note,
		Time:         now,
	}

	user.Bills = append(user.Bills, bill)

	err := db.Create(&bill).Error

	return bill.ID, err
}

// GetAllBills 获取用户账单列表
func GetAllBills(db *gorm.DB, userID uint64) ([]model.Bill, error) {
	var billList []model.Bill
	result := db.Where("user_id = ?", userID).Find(&billList)

	return billList, result.Error
}

// GetBillsByType 根据类型获取用户账单
func GetBillsByType(db *gorm.DB, userID uint64, moneyType string) ([]model.Bill, error) {
	var billList []model.Bill
	result := db.Find(&billList, "user_id = ? AND type = ?", userID, moneyType)

	return billList, result.Error
}

// DeleteBillByID 根据账单id删除账单记录
func DeleteBillByID(db *gorm.DB, id uint64) (float64, float64, error) {
	bill := new(model.Bill)
	err := db.Where("id = ?", id).First(bill).Error
	if err != nil {
		return 0, 0, err
	}
	err = db.Delete(&bill).Error
	return bill.Receipt, bill.Disbursement, err
}

// DeleteBills 删除列表中所有账单
func DeleteBills(db *gorm.DB, billList []model.Bill) error {
	for _, bill := range billList {
		err := db.Delete(&bill).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdateBillByID 修改用户收入/支出
func UpdateBillByID(db *gorm.DB, id uint64, receipt, disbursement float64, moneyType, note string) (float64, float64, error) {
	bill := new(model.Bill)
	err := db.Where("id = ?", id).First(bill).Error
	if err != nil {
		return 0, 0, err
	}
	oriReceipt := bill.Receipt
	oriDisbursement := bill.Disbursement
	err = db.Model(&bill).Updates(
		map[string]interface{}{
			"receipt":      receipt,
			"disbursement": disbursement,
			"type":         moneyType,
			"note":         note,
		}).Error

	return oriReceipt, oriDisbursement, err
}
