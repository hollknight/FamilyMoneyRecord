package stock

import (
	"FamilyMoneyRecord/database"
	"fmt"
	"testing"
)

func TestAddStock(t *testing.T) {
	db, _ := database.InitDB()
	stock, err := AddStock(db, 1, "sz000040", 20, 30)
	if err != nil {
		t.Errorf("添加股票信息失败：%s", err)
	}
	fmt.Println(stock)
}

func TestGetAllStocks(t *testing.T) {
	db, _ := database.InitDB()
	stockList, err := GetAllStocks(db)
	if err != nil {
		t.Errorf("获取股票持仓信息失败：%s", err)
	}
	fmt.Println(stockList)
}

func TestGetStock(t *testing.T) {
	db, _ := database.InitDB()
	stock, err := GetStock(db, 1, "sz000040")
	if err != nil {
		t.Errorf("获取股票持仓信息失败：%s", err)
	}
	fmt.Println(stock)
}

func TestGetStockByID(t *testing.T) {
	db, _ := database.InitDB()
	stock, err := GetStockByID(db, 1)
	if err != nil {
		t.Errorf("获取股票持仓信息失败：%s", err)
	}
	fmt.Println(stock)
}

func TestGetStocksByAccountID(t *testing.T) {
	db, _ := database.InitDB()
	stockList, err := GetStocksByAccountID(db, 1)
	if err != nil {
		t.Errorf("获取股票持仓信息失败：%s", err)
	}
	fmt.Println(stockList)
}

func TestUpdateStock(t *testing.T) {
	db, _ := database.InitDB()
	stock, err := GetStockByID(db, 1)
	if err != nil {
		t.Errorf("获取股票持仓信息失败：%s", err)
	}
	err = UpdateStock(db, stock, 40, 40)
	if err != nil {
		t.Errorf("更新股票信息失败：%s", err)
	}
}
