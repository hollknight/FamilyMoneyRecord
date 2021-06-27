package account

import (
	"FamilyMoneyRecord/database"
	"fmt"
	"testing"
)

func TestAddAccount(t *testing.T) {
	db, _ := database.InitDB()
	err := AddAccount(db, 1)
	if err != nil {
		t.Errorf("添加证券账户错误：%s", err)
	}
}

func TestUpdateAccountProfit(t *testing.T) {
	db, _ := database.InitDB()
	err := UpdateAccountProfit(db, 3, 30)
	if err != nil {
		t.Errorf("修改盈利出错：%s", err)
	}
}

func TestGetAccountsByUserID(t *testing.T) {
	db, _ := database.InitDB()
	accounts, err := GetAccountsByUserID(db, 1)
	if err != nil {
		t.Errorf("获取证券账户列表出错：%s", err)
	}
	fmt.Println(accounts)
}

func TestDeleteAccount(t *testing.T) {
	db, _ := database.InitDB()
	err := DeleteAccount(db, 3)
	if err != nil {
		t.Errorf("删除证券账户失败：%s", err)
	}
}
