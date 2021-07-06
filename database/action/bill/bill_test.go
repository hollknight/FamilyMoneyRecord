package bill

import (
	"FamilyMoneyRecord/database"
	"FamilyMoneyRecord/database/action/user"
	"fmt"
	"testing"
)

func TestAddBill(t *testing.T) {
	db, _ := database.InitDB()
	u, err := user.GetUserByUsername(db, "1234567891")
	if err != nil {
		t.Errorf("获取用户信息失败：%s", err)
	}
	id, err := AddBill(db, u, 0, 100, "衣食住行", "午餐")
	if err != nil {
		t.Errorf("添加账单失败：%s", err)
	}
	fmt.Println(id)
}

func TestGetBillsByUserID(t *testing.T) {
	db, _ := database.InitDB()
	billList, err := GetAllBillsByUserID(db, 4)
	if err != nil {
		t.Errorf("查看账单失败：%s", err)
	}
	fmt.Println(billList)
}

func TestDeleteBillByID(t *testing.T) {
	db, _ := database.InitDB()
	receipt, disbursement, err := DeleteBillByID(db, 3)
	if err != nil {
		t.Errorf("删除账单失败：%s", err)
	}
	fmt.Println(receipt, disbursement)
}

func TestGetAllBills(t *testing.T) {
	db, _ := database.InitDB()
	billList, err := GetAllBills(db)
	if err != nil {
		t.Errorf("获取账单失败：%s", err)
	}
	fmt.Println(billList)
}

func TestGetAllBillsByUserID(t *testing.T) {
	db, _ := database.InitDB()
	billList, err := GetAllBillsByUserID(db, 1)
	if err != nil {
		t.Errorf("获取账单失败：%s", err)
	}
	fmt.Println(billList)
}

func TestGetBillsByType(t *testing.T) {
	db, _ := database.InitDB()
	billList, err := GetBillsByType(db, 1, "衣食住行")
	if err != nil {
		t.Errorf("获取账单失败：%s", err)
	}
	fmt.Println(billList)
}

func TestUpdateBillByID(t *testing.T) {
	db, _ := database.InitDB()
	receipt, disbursement, err := UpdateBillByID(db, 4, 99, 0, "衣食住行", "吃饭")
	if err != nil {
		t.Errorf("修改账单失败：%s", err)
	}
	fmt.Println(receipt, disbursement)
}
