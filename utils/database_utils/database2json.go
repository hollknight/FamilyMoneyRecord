package database_utils

import (
	"FamilyMoneyRecord/config"
	"FamilyMoneyRecord/database/action/account"
	"FamilyMoneyRecord/database/action/bill"
	"FamilyMoneyRecord/database/action/operation"
	"FamilyMoneyRecord/database/action/stock"
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/database/model"
	"FamilyMoneyRecord/utils/resource_utils"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"io/ioutil"
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

// Struct2JSON 将结构体储存为json文件
func Struct2JSON(dataStruct Database, saveName string) error {
	path := config.FolderBathURL + saveName + ".json"
	isExist, err := resource_utils.IsExist(path)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New("该名称备份已存在")
	}

	filePtr, err := os.Create(path)
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

// JSON2Struct 读取json文件并转换为结构体
func JSON2Struct(jsonName string) (*Database, error) {
	database := new(Database)
	path := config.FolderBathURL + jsonName + ".json"
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteJson, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(byteJson, &database)
	return database, nil
}
