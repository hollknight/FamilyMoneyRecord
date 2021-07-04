package user

import (
	"FamilyMoneyRecord/database/action/account"
	"FamilyMoneyRecord/database/action/bill"
	"FamilyMoneyRecord/database/model"
	"gorm.io/gorm"
)

//// User 用户表
//type User struct {
//	ID                 uint64 `gorm:"primary_key;auto_increment"`
//	Username           string `gorm:"type:varchar(100);not null;unique"`
//	Password           string `gorm:"type:char(32);not null"`
//	Name               string `gorm:"type:varchar(20);default:''"`
//	ReceiptSum         int    `gorm:"type:int;default:0"`
//	DisbursementSum    int    `gorm:"type:int;default:0"`
//	AdvanceConsumption int    `gorm:"type:int;default:0"`
//	Bills              []bill2.Bill
//	Accounts           []account2.Account
//}

// AddUser 添加用户
func AddUser(db *gorm.DB, username, password, name string) error {
	user := model.User{
		Username:        username,
		Password:        password,
		Name:            name,
		ReceiptSum:      0,
		DisbursementSum: 0,
	}

	err := db.Create(&user).Error

	return err
}

// AddUserByStruct 根据结构体添加用户
func AddUserByStruct(db *gorm.DB, user model.User) error {
	err := db.Create(&user).Error

	return err
}

// GetUserByUsername 根据用户账号查询用户
func GetUserByUsername(db *gorm.DB, username string) (model.User, error) {
	user := new(model.User)
	err := db.Where("username = ?", username).First(user).Error

	return *user, err
}

// GetAllUsers 获取所有用户信息
func GetAllUsers(db *gorm.DB) ([]model.User, error) {
	var userList []model.User
	err := db.Find(&userList).Error

	return userList, err
}

// UpdateUserName 修改用户姓名
func UpdateUserName(db *gorm.DB, username, name string) error {
	err := db.Model(&model.User{}).Where("username = ?", username).Update("name", name).Error

	return err
}

// UpdateUserPassword 修改用户密码
func UpdateUserPassword(db *gorm.DB, user model.User, password string) error {
	err := db.Model(&user).Update("password", password).Error

	return err
}

// UpdateAdvanceConsumption 设置预消费金额
func UpdateAdvanceConsumption(db *gorm.DB, username string, advanceConsumption float64) error {
	err := db.Model(&model.User{}).Where("username = ?", username).Update("advance_consumption", advanceConsumption).Error

	return err
}

// DeleteUser 注销用户
func DeleteUser(db *gorm.DB, user model.User) error {
	// 删除所有与用户有关的账单
	billList, err := bill.GetAllBillsByUserID(db, user.ID)
	if err != nil {
		return err
	}
	err = bill.DeleteBills(db, billList)
	if err != nil {
		return err
	}

	// 删除所有与用户有关的证券账户
	accountList, err := account.GetAccountsByUserID(db, user.ID)
	if err != nil {
		return err
	}
	err = account.DeleteAccounts(db, accountList)
	if err != nil {
		return err
	}

	// 删除用户
	err = db.Delete(&user).Error
	//err := db.Select("Bills", "Accounts").Delete(&user).Error

	return err
}

// UpdateUserRSumAndDSum 增加收入金额
func UpdateUserRSumAndDSum(db *gorm.DB, user model.User, rSum, dSum float64) error {
	err := db.Model(&user).Updates(
		map[string]interface{}{
			"receipt_sum":      rSum,
			"disbursement_sum": dSum,
		}).Error

	return err
}

// PlusUserDisbursementSum 增加支出金额
func PlusUserDisbursementSum(db *gorm.DB, user model.User, money float64) error {
	err := db.Model(&user).Update("disbursement_sum", user.DisbursementSum+money).Error

	return err
}
