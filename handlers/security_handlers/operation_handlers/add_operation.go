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

type AddRequest struct {
	Token      string  `json:"token"`
	AccountID  uint64  `json:"accountID"`
	Code       string  `json:"code"`
	BuyNum     int     `json:"buyNum"`
	SaleNum    int     `json:"saleNum"`
	SharePrice float64 `json:"sharePrice"`
}

type AddResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (res *AddResponse) setAddResponse(code int, msg string) {
	res.Code = code
	res.Msg = msg
}

// AddOperation 添加账单
func AddOperation(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(AddRequest)
		response := new(AddResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setAddResponse(-1, "请检查传入参数是否缺失")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token := request.Token
		username, err := utils.ParseToken(token)
		if err != nil {
			response.setAddResponse(-2, "登录已过期，请重新登录")
			c.JSON(http.StatusOK, response)
			return
		}

		_, err = user.GetUserByUsername(db, username)
		if err != nil {
			response.setAddResponse(-3, "未查询到该账号")
			c.JSON(http.StatusOK, response)
			return
		}

		accountID := request.AccountID
		code := request.Code
		buyNum := request.BuyNum
		saleNum := request.SaleNum
		sharePrice := request.SharePrice
		s, err := stock.GetStock(db, accountID, code)
		if err != nil {
			if saleNum == 0 {
				s, err = stock.AddStock(db, accountID, code, 0, 0)
				if err != nil {
					response.setAddResponse(-4, "添加记录时出错，请稍后再试")
					c.JSON(http.StatusOK, response)
					return
				}
			} else {
				response.setAddResponse(-5, "股票持有股数不能小于0")
				c.JSON(http.StatusOK, response)
				return
			}
		}

		positionNum := s.PositionNum - saleNum + buyNum
		profit := s.Profit - float64(buyNum)*sharePrice + float64(saleNum)*sharePrice
		if positionNum < 0 {
			response.setAddResponse(-6, "股票持有股数不能小于0")
			c.JSON(http.StatusOK, response)
			return
		}

		err = db.Transaction(func(tx *gorm.DB) error {
			_, txErr := operation.AddOperation(tx, accountID, s.ID, sharePrice, buyNum, saleNum)
			if txErr != nil {
				return txErr
			}

			txErr = stock.UpdateStock(tx, s, positionNum, profit)
			if txErr != nil {
				return txErr
			}

			response.setAddResponse(0, "记录添加成功")
			c.JSON(http.StatusOK, response)
			return nil
		})
		if err != nil {
			response.setAddResponse(-4, "添加记录时发生错误")
			c.JSON(http.StatusOK, response)
		}
	}
}
