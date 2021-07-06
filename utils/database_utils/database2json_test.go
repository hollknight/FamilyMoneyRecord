package database_utils

import (
	database2 "FamilyMoneyRecord/database"
	"FamilyMoneyRecord/database/action/user"
	"fmt"
	"testing"
)

func TestStruct2JSON(t *testing.T) {
	db, _ := database2.InitDB()
	database, _ := SaveDatabase(db)
	err := Struct2JSON(database, "test1")
	if err != nil {
		t.Errorf("转换时发生错误：%s", err)
	}
}

func TestJSON2Struct(t *testing.T) {
	db, _ := database2.InitDB()
	database, err := JSON2Struct("test")
	if err != nil {
		t.Errorf("转换时发生时错误：%s", err)
	}
	user.AddUserByStruct(db, database.Users[0])
	fmt.Println(database.Users[0])
}
