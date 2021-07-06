package operation

import (
	"FamilyMoneyRecord/database"
	"fmt"
	"testing"
)

func TestAddOperation(t *testing.T) {
	db, _ := database.InitDB()
	id, err := AddOperation(db, 1, 1, 30, 20, 0)
	if err != nil {
		t.Errorf("添加股票交易记录失败：%s", err)
	}
	fmt.Println(id)
}

func TestGetAllOperations(t *testing.T) {
	db, _ := database.InitDB()
	operationList, err := GetAllOperations(db)
	if err != nil {
		t.Errorf("获取股票交易记录失败：%s", err)
	}
	fmt.Println(operationList)
}

func TestGetAllOperationsByAccountID(t *testing.T) {
	db, _ := database.InitDB()
	operationList, err := GetAllOperationsByAccountID(db, 1)
	if err != nil {
		t.Errorf("获取股票交易记录失败：%s", err)
	}
	fmt.Println(operationList)
}

func TestGetAllOperationsByStockID(t *testing.T) {
	db, _ := database.InitDB()
	operationList, err := GetAllOperationsByStockID(db, 1)
	if err != nil {
		t.Errorf("获取股票交易记录失败：%s", err)
	}
	fmt.Println(operationList)
}

func TestGetOperationByID(t *testing.T) {
	db, _ := database.InitDB()
	operation, err := GetOperationByID(db, 1)
	if err != nil {
		t.Errorf("获取股票交易记录失败：%s", err)
	}
	fmt.Println(operation)
}

func TestDeleteOperationByID(t *testing.T) {
	db, _ := database.InitDB()
	buyNum, SaleNum, price, err := DeleteOperationByID(db, 1)
	if err != nil {
		t.Errorf("删除股票交易记录失败：%s", err)
	}
	fmt.Println(buyNum, SaleNum, price)
}

func TestUpdateOperationByID(t *testing.T) {
	db, _ := database.InitDB()
	err := UpdateOperationByID(db, 1, 20, 0, 40)
	if err != nil {
		t.Errorf("更新股票交易记录失败：%s", err)
	}
}
