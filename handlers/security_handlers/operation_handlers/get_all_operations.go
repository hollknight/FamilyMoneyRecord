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

type AllRequest struct {
	Token     string `json:"token" binding:"required"`
	AccountID uint64 `json:"accountID" binding:"required"`
}

type AllResponse struct {
	Data AllData `json:"data"`
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
}

type AllData struct {
	Operations []AllOperation `json:"records"`
}

type AllOperation struct {
	Name       string  `json:"name"`
	Code       string  `json:"code"`
	BuyNum     int     `json:"buyNum"`
	SaleNum    int     `json:"saleNum"`
	SharePrice float64 `json:"sharePrice"`
	Time       string  `json:"time"`
}

func (res *AllResponse) setAllResponse(code int, msg string, records []AllOperation) {
	res.Data.Operations = records
	res.Code = code
	res.Msg = msg
}

// GetAllOperations 获取所有账单
func GetAllOperations(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(AllRequest)
		response := new(AllResponse)

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
		operationList, err := operation.GetAllOperationsByAccountID(db, accountID)
		if err != nil {
			response.setAllResponse(-4, "获取时发生错误，请稍后再试", nil)
			c.JSON(http.StatusOK, response)
			return
		}
		var records []AllOperation
		for _, operationRecord := range operationList {
			s, err := stock.GetStockByID(db, operationRecord.StockID)
			if err != nil {
				response.setAllResponse(-5, "获取时发生错误，请稍后再试", nil)
				c.JSON(http.StatusOK, response)
				return
			}
			name, _, err := stock_info_utils.GetStockInfo(s.Code)
			if err != nil {
				response.setAllResponse(-6, "获取时发生错误，请稍后再试", nil)
				c.JSON(http.StatusOK, response)
				return
			}
			timeRecord := time.Unix(operationRecord.Time.Unix(), 0).Format("2006-01-02 15:04:05")
			record := AllOperation{
				Name:       name,
				Code:       s.Code,
				BuyNum:     operationRecord.BuyNum,
				SaleNum:    operationRecord.SaleNum,
				SharePrice: operationRecord.SharePrice,
				Time:       timeRecord,
			}
			records = append(records, record)
		}

		response.setAllResponse(0, "查询成功", records)
		c.JSON(http.StatusOK, response)
	}
}
