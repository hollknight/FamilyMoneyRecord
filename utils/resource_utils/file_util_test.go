package resource_utils

import (
	"fmt"
	"testing"
)

func TestGetJSONName(t *testing.T) {
	names, err := GetJSONName()
	if err != nil {
		t.Errorf("获取文件名错误：%s", err)
	}
	fmt.Println(names)
}
