package utils

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func GetStockPrice(code string) (float64, error) {
	url := "http://hq.sinajs.cn/list=" + code
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	arr := strings.Split(string(s), ",")
	price, err := strconv.ParseFloat(arr[3], 64)
	return price, err
}
