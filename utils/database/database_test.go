package database

import (
	"FamilyMoneyRecord/database"
	"testing"
)

func TestDropAllTables(t *testing.T) {
	db, _ := database.InitDB()
	DropAllTables(db)
}

func TestCreateTables(t *testing.T) {
	db, _ := database.InitDB()
	DropAllTables(db)
	err := CreateTables(db)
	if err != nil {
		t.Errorf("创建表单时发生错误：%s", err)
	}
}
