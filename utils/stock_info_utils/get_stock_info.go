package stock_info_utils

import (
	"errors"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func GetStockInfo(code string) (string, float64, error) {
	url := "http://hq.sinajs.cn/list=" + code
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	arr := strings.Split(string(s), ",")
	if len(arr) <= 1 {
		return "", 0, errors.New("错误的股票代码")
	}
	ar := strings.Split(arr[0], "\"")
	nameGB18030 := ar[1]
	nameUTF8, _ := simplifiedchinese.GB18030.NewDecoder().Bytes([]byte(nameGB18030))
	price, err := strconv.ParseFloat(arr[3], 64)
	return string(nameUTF8), price, err
}
