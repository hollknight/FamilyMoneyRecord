package bill_handlers

import (
	"FamilyMoneyRecord/database/action/bill"
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type DeleteRequest struct {
	Token string `json:"token" binding:"required"`
	ID    uint64 `json:"id" binding:"required"`
}

type DeleteResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (res *DeleteResponse) setDeleteResponse(code int, msg string) {
	res.Code = code
	res.Msg = msg
}

// DeleteBill 删除账单
func DeleteBill(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(DeleteRequest)
		response := new(DeleteResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setDeleteResponse(-1, "请检查传入参数是否缺失")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token := request.Token
		username, err := utils.ParseToken(token)
		if err != nil {
			response.setDeleteResponse(-2, "登录已过期，请重新登录")
			c.JSON(http.StatusOK, response)
			return
		}

		u, err := user.GetUserByUsername(db, username)
		if err != nil {
			response.setDeleteResponse(-3, "获取用户失败，请重新登录")
			c.JSON(http.StatusOK, response)
			return
		}

		id := request.ID
		err = db.Transaction(func(tx *gorm.DB) error {
			receipt, disbursement, txErr := bill.DeleteBillByID(tx, id)
			if txErr != nil {
				return txErr
			}
			txErr = user.UpdateUserRSumAndDSum(tx, u, u.ReceiptSum-receipt, u.DisbursementSum-disbursement)
			if txErr != nil {
				return txErr
			}
			response.setDeleteResponse(0, "删除账单成功")
			c.JSON(http.StatusOK, response)
			return nil
		})
		if err != nil {
			response.setDeleteResponse(-4, "删除时发生错误，请稍后再试")
			c.JSON(http.StatusOK, response)
		}
	}
}
