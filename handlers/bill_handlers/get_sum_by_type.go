package bill_handlers

import (
	"FamilyMoneyRecord/database/action/bill"
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type SumRequest struct {
	Token string `json:"token" binding:"required"`
}

type SumResponse struct {
	Data SumData `json:"data"`
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
}

type SumData struct {
	Receipt      []float64 `json:"receipt"`
	Disbursement []float64 `json:"disbursement"`
}

func (res *SumResponse) setSumResponse(code int, msg string, receipt, disbursement []float64) {
	res.Data = SumData{
		Receipt:      receipt,
		Disbursement: disbursement,
	}
	res.Code = code
	res.Msg = msg
}

// GetSumByType 删除账单
func GetSumByType(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(SumRequest)
		response := new(SumResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setSumResponse(-1, "请检查传入参数是否缺失", nil, nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token := request.Token
		username, err := utils.ParseToken(token)
		if err != nil {
			response.setSumResponse(-2, "登录已过期，请重新登录", nil, nil)
			c.JSON(http.StatusOK, response)
			return
		}

		u, err := user.GetUserByUsername(db, username)
		if err != nil {
			response.setSumResponse(-3, "获取用户失败，请重新登录", nil, nil)
			c.JSON(http.StatusOK, response)
			return
		}

		var receipts []float64
		var disbursements []float64

		receiptType := []string{"工资", "股票", "分红", "奖金"}
		disbursementType := []string{"税收", "衣食住行", "医疗", "其他"}

		for _, rType := range receiptType {
			var typeReceipt float64
			typeReceipt = 0
			billList, err := bill.GetBillsByType(db, u.ID, rType)
			if err != nil {
				response.setSumResponse(-4, "获取时发生错误，请稍后再试", nil, nil)
				c.JSON(http.StatusOK, response)
			}
			for _, bill := range billList {
				typeReceipt += bill.Receipt
			}
			receipts = append(receipts, typeReceipt)
		}

		for _, dType := range disbursementType {
			var typeDisbursement float64
			typeDisbursement = 0
			billList, err := bill.GetBillsByType(db, u.ID, dType)
			if err != nil {
				response.setSumResponse(-5, "获取时发生错误，请稍后再试", nil, nil)
				c.JSON(http.StatusOK, response)
			}
			for _, bill := range billList {
				typeDisbursement += bill.Disbursement
			}
			disbursements = append(disbursements, typeDisbursement)
		}

		response.setSumResponse(0, "查询成功", receipts, disbursements)
		c.JSON(http.StatusOK, response)
	}
}
