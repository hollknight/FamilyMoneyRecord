package database

import (
	"fmt"
	"testing"
)

func TestInitDB(t *testing.T) {
	_, err := InitDB()
	fmt.Println(err)
}
