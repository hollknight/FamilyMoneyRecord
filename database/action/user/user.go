package user

import (
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

	result := db.Create(&user)

	return result.Error
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
	result := db.Find(&userList)

	return userList, result.Error
}

// UpdateUserName 修改用户姓名
func UpdateUserName(db *gorm.DB, username, name string) error {
	err := db.Model(&model.User{}).Where("username = ?", username).Update("name", name).Error

	return err
}

// UpdateUserPassword 修改用户密码
func UpdateUserPassword(db *gorm.DB, username, password string) error {
	err := db.Model(&model.User{}).Where("username = ?", username).Update("password", password).Error

	return err
}

// UpdateAdvanceConsumption 设置预消费金额
func UpdateAdvanceConsumption(db *gorm.DB, username string, advanceConsumption int) error {
	err := db.Model(&model.User{}).Where("username = ?", username).Update("advance_consumption", advanceConsumption).Error

	return err
}

// DeleteUser 注销用户
func DeleteUser(db *gorm.DB, user model.User) error {
	err := db.Delete(&user).Error

	return err
}

// PlusUserReceiptSum 增加收入金额
func PlusUserReceiptSum(db *gorm.DB, user model.User, money int) error {
	err := db.Model(&user).Update("receipt_sum", user.ReceiptSum+money).Error

	return err
}

// PlusUserDisbursementSum 增加支出金额
func PlusUserDisbursementSum(db *gorm.DB, user model.User, money int) error {
	err := db.Model(&user).Update("disbursement_sum", user.DisbursementSum+money).Error

	return err
}
