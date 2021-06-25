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
func AddBill(db *gorm.DB, user model.User, receipt, disbursement int, moneyType string) error {
	bill := model.Bill{
		UserID:       user.ID,
		Receipt:      receipt,
		Disbursement: disbursement,
		Type:         moneyType,
		Time:         time.Now(),
	}

	user.Bills = append(user.Bills, bill)

	result := db.Create(&bill)

	return result.Error
}
