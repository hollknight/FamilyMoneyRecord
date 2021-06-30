package bill_handlers

import (
	"FamilyMoneyRecord/database/action/bill"
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type AddRequest struct {
	Token        string  `json:"token"`
	Receipt      float64 `json:"receipt"`
	Disbursement float64 `json:"disbursement"`
	Type         string  `json:"type"`
	Note         string  `json:"note"`
}

type AddResponse struct {
	Data AddData `json:"data"`
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
}

type AddData struct {
	ID uint64 `json:"id"`
}

func (res *AddResponse) setAddResponse(code int, msg string, id uint64) {
	res.Data.ID = id
	res.Code = code
	res.Msg = msg
}

// AddBill 添加账单
func AddBill(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(AddRequest)
		response := new(AddResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setAddResponse(-1, "请检查传入参数是否缺失", 0)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token := request.Token
		username, err := utils.ParseToken(token)
		if err != nil {
			response.setAddResponse(-2, "登录已过期，请重新登录", 0)
			c.JSON(http.StatusOK, response)
			return
		}

		u, err := user.GetUserByUsername(db, username)
		if err != nil {
			response.setAddResponse(-3, "未查询到该账号", 0)
			c.JSON(http.StatusOK, response)
			return
		}

		receipt := request.Receipt
		disbursement := request.Disbursement
		moneyType := request.Type
		note := request.Note

		//id, err := bill.AddBill(db, u, receipt, disbursement, moneyType, note)
		//if err != nil {
		//	response.setAddResponse(-4, "添加账单出错，请稍后再试", 0)
		//	c.JSON(http.StatusOK, response)
		//	return
		//}
		//err = user.UpdateUserRSumAndDSum(db, u, u.ReceiptSum+receipt, u.DisbursementSum+disbursement)
		//if err != nil {
		//	bill.DeleteBillByID(db, id)
		//	response.setAddResponse(-5, "修改金额过程中发生错误", 0)
		//	c.JSON(http.StatusOK, response)
		//	return
		//}

		err = db.Transaction(func(tx *gorm.DB) error {
			id, txErr := bill.AddBill(tx, u, receipt, disbursement, moneyType, note)
			if txErr != nil {
				return txErr
			}
			txErr = user.UpdateUserRSumAndDSum(tx, u, u.ReceiptSum+receipt, u.DisbursementSum+disbursement)
			if txErr != nil {
				return txErr
			}

			response.setAddResponse(0, "账单添加成功", id)
			c.JSON(http.StatusOK, response)
			return nil
		})
		if err != nil {
			response.setAddResponse(-4, "添加账单时发生错误", 0)
			c.JSON(http.StatusOK, response)
		}
	}
}
