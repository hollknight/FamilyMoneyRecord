package user

import (
	"FamilyMoneyRecord/database"
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {
	db, _ := database.InitDB()
	err := AddUser(db, "1234567892", "12345678901234567890123456789012", "name")
	if err != nil {
		t.Errorf("添加用户失败：%s", err)
	}
}

func TestGetUserByUsername(t *testing.T) {
	db, _ := database.InitDB()
	user, err := GetUserByUsername(db, "1234567890")
	if err != nil {
		t.Errorf("获取用户信息失败：%s", err)
	}
	fmt.Println(user)
}

func TestGetUsersByLikeUsername(t *testing.T) {
	db, _ := database.InitDB()
	users, err := GetUsersByLikeUsername(db, "admin")
	if err != nil {
		t.Errorf("获取用户信息失败：%s", err)
	}
	fmt.Println(users)
}

func TestGetAllUsers(t *testing.T) {
	db, _ := database.InitDB()
	userList, err := GetAllUsers(db)
	if err != nil {
		t.Errorf("获取所有用户信息失败：%s", err)
	}
	fmt.Println(userList)
}

func TestUpdateUserName(t *testing.T) {
	db, _ := database.InitDB()
	err := UpdateUserName(db, "1234567890", "update")
	if err != nil {
		t.Errorf("修改用户姓名失败：%s", err)
	}
}

func TestUpdateUserPassword(t *testing.T) {
	//db, _ := database.InitDB()
	//err := UpdateUserPassword(db, "1234567890", "12345678901234567890123456789099")
	//if err != nil {
	//	t.Errorf("修改用户密码失败：%s", err)
	//}
}

func TestUpdateAdvanceConsumption(t *testing.T) {
	db, _ := database.InitDB()
	err := UpdateAdvanceConsumption(db, "1234567890", 20)
	if err != nil {
		t.Errorf("修改设置用户预消费金额失败：%s", err)
	}
}

func TestDeleteUser(t *testing.T) {
	db, _ := database.InitDB()
	user, err := GetUserByUsername(db, "1234567890")
	if err != nil {
		t.Errorf("查询用户失败：%s", err)
	}
	err = DeleteUser(db, user)
	if err != nil {
		t.Errorf("删除用户失败：%s", err)
	}
}

func TestPlusUserReceiptSum(t *testing.T) {
	db, _ := database.InitDB()
	user, err := GetUserByUsername(db, "1234567890")
	if err != nil {
		t.Errorf("查询用户失败：%s", err)
	}
	err = UpdateUserRSumAndDSum(db, user, 5, 5)
	if err != nil {
		t.Errorf("用户更新收入失败：%s", err)
	}
}

func TestPlusUserDisbursementSum(t *testing.T) {
	db, _ := database.InitDB()
	user, err := GetUserByUsername(db, "1234567890")
	if err != nil {
		t.Errorf("查询用户失败：%s", err)
	}
	err = PlusUserDisbursementSum(db, user, 5)
	if err != nil {
		t.Errorf("用户更新支出失败：%s", err)
	}
}

func TestUpdateUserRSumAndDSum(t *testing.T) {
	db, _ := database.InitDB()
	user, err := GetUserByUsername(db, "ai")
	if err != nil {
		t.Errorf("查询用户失败：%s", err)
	}
	err = UpdateUserRSumAndDSum(db, user, 0, 0)
	if err != nil {
		t.Errorf("用户更新收支信息失败：%s", err)
	}
}
