package account

import (
	"FamilyMoneyRecord/database"
	"fmt"
	"testing"
)

func TestAddAccount(t *testing.T) {
	db, _ := database.InitDB()
	id, err := AddAccount(db, 1)
	if err != nil {
		t.Errorf("添加证券账户错误：%s", err)
	}
	fmt.Println(id)
}

func TestGetAccountsByUserID(t *testing.T) {
	db, _ := database.InitDB()
	accounts, err := GetAccountsByUserID(db, 1)
	if err != nil {
		t.Errorf("获取证券账户列表出错：%s", err)
	}
	fmt.Println(accounts)
}

func TestGetAllAccounts(t *testing.T) {
	db, _ := database.InitDB()
	accountList, err := GetAllAccounts(db)
	if err != nil {
		t.Errorf("获取证券账户列表出错：%s", err)
	}
	fmt.Println(accountList)
}

func TestDeleteAccount(t *testing.T) {
	db, _ := database.InitDB()
	err := DeleteAccount(db, 3)
	if err != nil {
		t.Errorf("删除证券账户失败：%s", err)
	}
}
