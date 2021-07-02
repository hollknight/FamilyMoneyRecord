package database_utils

import (
	"FamilyMoneyRecord/config"
	"FamilyMoneyRecord/database/action/account"
	"FamilyMoneyRecord/database/action/bill"
	"FamilyMoneyRecord/database/action/operation"
	"FamilyMoneyRecord/database/action/stock"
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/database/model"
	"encoding/json"
	"gorm.io/gorm"
	"os"
)

type Database struct {
	Users      []model.User      `json:"users"`
	Bills      []model.Bill      `json:"bills"`
	Accounts   []model.Account   `json:"accounts"`
	Stocks     []model.Stock     `json:"stocks"`
	Operations []model.Operation `json:"operations"`
}

// SaveDatabase 读取数据库中所有数据并存储
func SaveDatabase(db *gorm.DB) (Database, error) {
	users, err := user.GetAllUsers(db)
	bills, err := bill.GetAllBills(db)
	accounts, err := account.GetAllAccounts(db)
	stocks, err := stock.GetAllStocks(db)
	operations, err := operation.GetAllOperations(db)

	database := Database{
		Users:      users,
		Bills:      bills,
		Accounts:   accounts,
		Stocks:     stocks,
		Operations: operations,
	}

	return database, err
}

// Struct2json 将结构体储存为json文件
func Struct2json(dataStruct Database, saveName string) error {
	name := config.FolderBathURL + saveName + ".json"
	filePtr, err := os.Create(name)
	if err != nil {
		return err
	}

	defer filePtr.Close()

	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(dataStruct)
	if err != nil {
		return err
	}

	return nil
}
