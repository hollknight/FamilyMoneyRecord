package stock_handlers

import (
	"FamilyMoneyRecord/database/action/stock"
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"FamilyMoneyRecord/utils/stock_info_utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type StockRequest struct {
	Token     string `json:"token" binding:"required"`
	AccountID uint64 `json:"accountID" binding:"required"`
}

type StockResponse struct {
	Data StockData `json:"data"`
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
}

type StockData struct {
	Stocks []StockRecord `json:"stocks"`
}

type StockRecord struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	Code        string  `json:"code"`
	PositionNum int     `json:"positionNum"`
	Price       float64 `json:"price"`
	Profit      float64 `json:"profit"`
}

func (res *StockResponse) setAllResponse(code int, msg string, records []StockRecord) {
	res.Data.Stocks = records
	res.Code = code
	res.Msg = msg
}

// GetAllStocks 获取所有证券账户
func GetAllStocks(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(StockRequest)
		response := new(StockResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setAllResponse(-1, "请检查传入参数是否缺失", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token := request.Token
		username, err := utils.ParseToken(token)
		if err != nil {
			response.setAllResponse(-2, "登录已过期，请重新登录", nil)
			c.JSON(http.StatusOK, response)
			return
		}

		_, err = user.GetUserByUsername(db, username)
		if err != nil {
			response.setAllResponse(-3, "获取用户失败，请重新登录", nil)
			c.JSON(http.StatusOK, response)
			return
		}

		accountID := request.AccountID
		stockList, err := stock.GetStocksByAccountID(db, accountID)
		if err != nil {
			response.setAllResponse(-4, "获取时发生错误，请稍后再试", nil)
			c.JSON(http.StatusOK, response)
			return
		}
		var records []StockRecord
		for _, stockRecord := range stockList {
			name, price, err := stock_info_utils.GetStockInfo(stockRecord.Code)
			if err != nil {
				response.setAllResponse(-5, "获取价格时发生错误，请稍后再试", nil)
				c.JSON(http.StatusOK, response)
				return
			}
			profit := stockRecord.Profit + float64(stockRecord.PositionNum)*price
			record := StockRecord{
				ID:          stockRecord.ID,
				Name:        name,
				Code:        stockRecord.Code,
				PositionNum: stockRecord.PositionNum,
				Price:       price,
				Profit:      profit,
			}
			records = append(records, record)
		}

		response.setAllResponse(0, "查询成功", records)
		c.JSON(http.StatusOK, response)
	}
}
