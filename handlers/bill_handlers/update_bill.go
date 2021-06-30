package bill_handlers

import (
	"FamilyMoneyRecord/database/action/bill"
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type UpdateRequest struct {
	Token        string  `json:"token" binding:"required"`
	ID           uint64  `json:"id" binding:"required"`
	Receipt      float64 `json:"receipt" binding:"required"`
	Disbursement float64 `json:"disbursement" binding:"required"`
	Type         string  `json:"type" binding:"required"`
	Note         string  `json:"note" binding:"required"`
}

type UpdateResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (res *UpdateResponse) setUpdateResponse(code int, msg string) {
	res.Code = code
	res.Msg = msg
}

// UpdateBill 删除账单
func UpdateBill(db *gorm.DB) func(c *gin.Context) {
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

		u, err := user.GetUserByUsername(db, username)
		if err != nil {
			response.setUpdateResponse(-3, "获取用户失败，请重新登录")
			c.JSON(http.StatusOK, response)
			return
		}

		id := request.ID
		receipt := request.Receipt
		disbursement := request.Disbursement
		moneyType := request.Type
		note := request.Note
		err = db.Transaction(func(tx *gorm.DB) error {
			oriReceipt, oriDisbursement, txErr := bill.UpdateBillByID(tx, id, receipt, disbursement, moneyType, note)
			if txErr != nil {
				return txErr
			}
			txErr = user.UpdateUserRSumAndDSum(tx, u, u.ReceiptSum+receipt-oriReceipt, u.DisbursementSum+disbursement-oriDisbursement)
			if txErr != nil {
				return txErr
			}
			response.setUpdateResponse(0, "修改账单成功")
			c.JSON(http.StatusOK, response)
			return nil
		})
		if err != nil {
			response.setUpdateResponse(-4, "修改时发生错误，请稍后再试")
			c.JSON(http.StatusOK, response)
		}
	}
}
