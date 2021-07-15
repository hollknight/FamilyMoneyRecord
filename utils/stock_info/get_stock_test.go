package stock_info

import (
	"fmt"
	"testing"
)

func TestGetStockInfo(t *testing.T) {
	name, price, err := GetStockInfo("sz002307")
	if err != nil {
		t.Errorf("获取股票信息时发生错误：%s", err)
	}
	fmt.Println(name)
	fmt.Println(price)
}
