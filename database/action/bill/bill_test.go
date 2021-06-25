package bill

import (
	"FamilyMoneyRecord/database"
	"FamilyMoneyRecord/database/action/user"
	"fmt"
	"testing"
)

func TestAddBill(t *testing.T) {
	db, _ := database.InitDB()
	u, err := user.GetUserByUsername(db, "1234567890")
	if err != nil {
		t.Errorf("获取用户信息失败：%s", err)
	}
	id, err := AddBill(db, u, 20, 0, "吃饭")
	if err != nil {
		t.Errorf("添加账单失败：%s", err)
	}
	fmt.Println(id)
}
