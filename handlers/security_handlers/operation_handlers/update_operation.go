package operation_handlers

import (
	"FamilyMoneyRecord/database/action/operation"
	"FamilyMoneyRecord/database/action/stock"
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type UpdateRequest struct {
	Token      string  `json:"token"`
	ID         uint64  `json:"id"`
	AccountID  uint64  `json:"accountID"`
	Code       string  `json:"code"`
	BuyNum     int     `json:"buyNum"`
	SaleNum    int     `json:"saleNum"`
	SharePrice float64 `json:"sharePrice"`
}

type UpdateResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (res *UpdateResponse) setUpdateResponse(code int, msg string) {
	res.Code = code
	res.Msg = msg
}

// UpdateOperation 删除账单
func UpdateOperation(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(UpdateRequest)
		response := new(UpdateResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setUpdateResponse(-1, "请检查传入参数是否缺失")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token := request.Token
		username, err := utils.ParseToken(token)
		if err != nil {
			response.setUpdateResponse(-2, "登录已过期，请重新登录")
			c.JSON(http.StatusOK, response)
			return
		}

		_, err = user.GetUserByUsername(db, username)
		if err != nil {
			response.setUpdateResponse(-3, "获取用户失败，请重新登录")
			c.JSON(http.StatusOK, response)
			return
		}

		id := request.ID
		accountID := request.AccountID
		code := request.Code
		buyNum := request.BuyNum
		saleNum := request.SaleNum
		sharePrice := request.SharePrice
		s, err := stock.GetStock(db, accountID, code)
		if err != nil {
			response.setUpdateResponse(-4, "未查询到该股票")
			c.JSON(http.StatusOK, response)
			return
		}

		o, err := operation.GetOperationByID(db, id)
		if err != nil {
			response.setUpdateResponse(-5, "未查询到该交易记录")
			c.JSON(http.StatusOK, response)
			return
		}

		positionNum := s.PositionNum + buyNum - o.BuyNum - saleNum + o.SaleNum
		profit := s.Profit + float64(o.BuyNum)*o.SharePrice - float64(buyNum)*sharePrice + float64(o.SaleNum)*o.SharePrice - float64(saleNum)*sharePrice
		if positionNum < 0 {
			response.setUpdateResponse(-6, "修改交易记录失败，持有股数不能为负数")
			c.JSON(http.StatusOK, response)
			return
		}

		err = db.Transaction(func(tx *gorm.DB) error {
			txErr := operation.UpdateOperationByID(tx, id, buyNum, saleNum, sharePrice)
			if txErr != nil {
				return txErr
			}
			txErr = stock.UpdateStock(tx, s, positionNum, profit)
			if txErr != nil {
				return txErr
			}
			response.setUpdateResponse(0, "修改交易记录成功")
			c.JSON(http.StatusOK, response)
			return nil
		})
		if err != nil {
			response.setUpdateResponse(-7, "修改时发生错误，请稍后再试")
			c.JSON(http.StatusOK, response)
		}
	}
}
