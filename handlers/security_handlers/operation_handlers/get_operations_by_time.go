package operation_handlers

import (
	"FamilyMoneyRecord/database/action/operation"
	"FamilyMoneyRecord/database/action/stock"
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"FamilyMoneyRecord/utils/stock_info_utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type TimeRequest struct {
	Token     string `json:"token" binding:"required"`
	AccountID uint64 `json:"accountID" binding:"required"`
	BeginTime string `json:"beginTime" binding:"required"`
	EndTime   string `json:"endTime" binding:"required"`
}

type TimeResponse struct {
	Data TimeData `json:"data"`
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
}

type TimeData struct {
	Operations []TimeOperation `json:"records"`
}

type TimeOperation struct {
	ID         uint64  `json:"id"`
	Name       string  `json:"name"`
	Code       string  `json:"code"`
	BuyNum     int     `json:"buyNum"`
	SaleNum    int     `json:"saleNum"`
	SharePrice float64 `json:"sharePrice"`
	Time       string  `json:"time"`
}

func (res *TimeResponse) setTimeResponse(code int, msg string, records []TimeOperation) {
	res.Data.Operations = records
	res.Code = code
	res.Msg = msg
}

// GetOperationsByTime 获取所有账单
func GetOperationsByTime(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(TimeRequest)
		response := new(TimeResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setTimeResponse(-1, "请检查传入参数是否缺失", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token := request.Token
		username, err := utils.ParseToken(token)
		if err != nil {
			response.setTimeResponse(-2, "登录已过期，请重新登录", nil)
			c.JSON(http.StatusOK, response)
			return
		}

		_, err = user.GetUserByUsername(db, username)
		if err != nil {
			response.setTimeResponse(-3, "获取用户失败，请重新登录", nil)
			c.JSON(http.StatusOK, response)
			return
		}

		beginTimeStr := request.BeginTime
		beginTime, _ := time.Parse("2006-01-02 15:04:05", beginTimeStr)
		endTimeStr := request.EndTime
		endTime, _ := time.Parse("2006-01-02 15:04:05", endTimeStr)
		accountID := request.AccountID
		operationList, err := operation.GetAllOperationsByAccountID(db, accountID)
		if err != nil {
			response.setTimeResponse(-4, "获取时发生错误，请稍后再试", nil)
			c.JSON(http.StatusOK, response)
			return
		}
		var records []TimeOperation
		for _, operationRecord := range operationList {
			if beginTime.Before(operationRecord.Time) && endTime.After(operationRecord.Time) {
				s, err := stock.GetStockByID(db, operationRecord.StockID)
				if err != nil {
					response.setTimeResponse(-5, "获取时发生错误，请稍后再试", nil)
					c.JSON(http.StatusOK, response)
					return
				}
				name, _, err := stock_info_utils.GetStockInfo(s.Code)
				if err != nil {
					response.setTimeResponse(-6, "无效的股票代码", nil)
					c.JSON(http.StatusOK, response)
					return
				}
				timeRecord := time.Unix(operationRecord.Time.Unix(), 0).Format("2006-01-02 15:04:05")
				record := TimeOperation{
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
		}

		response.setTimeResponse(0, "查询成功", records)
		c.JSON(http.StatusOK, response)
	}
}
