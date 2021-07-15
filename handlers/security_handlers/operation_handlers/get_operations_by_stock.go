package operation_handlers

import (
	"FamilyMoneyRecord/database/action/operation"
	"FamilyMoneyRecord/database/action/stock"
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"FamilyMoneyRecord/utils/stock_info"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type StockRequest struct {
	Token     string `json:"token" binding:"required"`
	AccountID uint64 `json:"accountID" binding:"required"`
	Code      string `json:"code" binding:"required"`
}

type StockResponse struct {
	Data StockData `json:"data"`
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
}

type StockData struct {
	Operations []StockOperation `json:"records"`
}

type StockOperation struct {
	ID         uint64  `json:"id"`
	Name       string  `json:"name"`
	Code       string  `json:"code"`
	BuyNum     int     `json:"buyNum"`
	SaleNum    int     `json:"saleNum"`
	SharePrice float64 `json:"sharePrice"`
	Time       string  `json:"time"`
}

func (res *StockResponse) setStockResponse(code int, msg string, records []StockOperation) {
	res.Data.Operations = records
	res.Code = code
	res.Msg = msg
}

// GetOperationsByStock 获取所有账单
func GetOperationsByStock(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(StockRequest)
		response := new(StockResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setStockResponse(-1, "请检查传入参数是否缺失", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token := request.Token
		username, err := utils.ParseToken(token)
		if err != nil {
			response.setStockResponse(-2, "登录已过期，请重新登录", nil)
			c.JSON(http.StatusOK, response)
			return
		}

		_, err = user.GetUserByUsername(db, username)
		if err != nil {
			response.setStockResponse(-3, "获取用户失败，请重新登录", nil)
			c.JSON(http.StatusOK, response)
			return
		}

		accountID := request.AccountID
		code := request.Code
		s, err := stock.GetStock(db, accountID, code)
		if err != nil {
			response.setStockResponse(-4, "获取时发生错误，请稍后再试", nil)
			c.JSON(http.StatusOK, response)
			return
		}
		operationList, err := operation.GetAllOperationsByStockID(db, s.ID)
		if err != nil {
			response.setStockResponse(-5, "获取时发生错误，请稍后再试", nil)
			c.JSON(http.StatusOK, response)
			return
		}
		var records []StockOperation
		for _, operationRecord := range operationList {
			s, err := stock.GetStockByID(db, s.ID)
			if err != nil {
				response.setStockResponse(-6, "获取时发生错误，请稍后再试", nil)
				c.JSON(http.StatusOK, response)
				return
			}
			name, _, err := stock_info.GetStockInfo(s.Code)
			if err != nil {
				response.setStockResponse(-7, "无效的股票代码", nil)
				c.JSON(http.StatusOK, response)
				return
			}
			timeRecord := time.Unix(operationRecord.Time.Unix(), 0).Format("2006-01-02 15:04:05")
			record := StockOperation{
				ID:         operationRecord.ID,
				Name:       name,
				Code:       s.Code,
				BuyNum:     operationRecord.BuyNum,
				SaleNum:    operationRecord.SaleNum,
				SharePrice: operationRecord.SharePrice,
				Time:       timeRecord,
			}
			records = append(records, record)
		}

		response.setStockResponse(0, "查询成功", records)
		c.JSON(http.StatusOK, response)
	}
}
