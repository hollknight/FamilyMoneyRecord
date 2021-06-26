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
func AddBill(db *gorm.DB, user model.User, receipt, disbursement int, moneyType string) (uint64, error) {
	now := time.Now().Add(time.Hour * 8)

	bill := model.Bill{
		UserID:       user.ID,
		Receipt:      receipt,
		Disbursement: disbursement,
		Type:         moneyType,
		Time:         now,
	}

	user.Bills = append(user.Bills, bill)

	result := db.Create(&bill)

	return bill.ID, result.Error
}

// GetBillsByUserID 获取用户账单列表
func GetBillsByUserID(db *gorm.DB, userID uint64) ([]model.Bill, error) {
	var billList []model.Bill
	result := db.Where("user_id = ?", userID).Find(&billList)

	return billList, result.Error
}

// DeleteBillByID 根据账单id删除账单记录
func DeleteBillByID(db *gorm.DB, id uint64) (int, int, error) {
	bill := new(model.Bill)
	err := db.Where("id = ?", id).First(bill).Error
	if err != nil {
		return 0, 0, err
	}
	err = db.Delete(&bill).Error
	return bill.Receipt, bill.Disbursement, err
}

// UpdateBillByID 修改用户收入/支出
func UpdateBillByID(db *gorm.DB, id uint64, receipt, disbursement int) (int, int, error) {
	bill := new(model.Bill)
	err := db.Where("id = ?", id).First(bill).Error
	if err != nil {
		return 0, 0, err
	}
	err = db.Model(&bill).Updates(
		model.Bill{
			Receipt:      receipt,
			Disbursement: disbursement,
		}).Error

	return bill.Receipt, bill.Disbursement, err
}
