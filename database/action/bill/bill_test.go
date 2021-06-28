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
	id, err := AddBill(db, u, 0, 100, "吃饭")
	if err != nil {
		t.Errorf("添加账单失败：%s", err)
	}
	fmt.Println(id)
}

func TestGetBillsByUserID(t *testing.T) {
	db, _ := database.InitDB()
	billList, err := GetBillsByUserID(db, 4)
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

func TestUpdateBillByID(t *testing.T) {
	db, _ := database.InitDB()
	receipt, disbursement, err := UpdateBillByID(db, 4, 99, 0)
	if err != nil {
		t.Errorf("修改账单失败：%s", err)
	}
	fmt.Println(receipt, disbursement)
}
