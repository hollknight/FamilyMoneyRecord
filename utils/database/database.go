package database

import (
	"FamilyMoneyRecord/database/model"
	"gorm.io/gorm"
)

// DropAllTables 删除所有表
func DropAllTables(db *gorm.DB) {
	db.Exec("DROP TABLE operations")
	db.Exec("DROP TABLE stocks")
	db.Exec("DROP TABLE accounts")
	db.Exec("DROP TABLE bills")
	db.Exec("DROP TABLE users")
}

// CreateTables 初始化创建所有表
func CreateTables(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{}, &model.Bill{}, &model.Account{}, &model.Stock{}, &model.Operation{})
	return err
}
