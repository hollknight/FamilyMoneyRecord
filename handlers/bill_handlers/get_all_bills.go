package bill_handlers

import (
	"FamilyMoneyRecord/database/action/bill"
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type AllRequest struct {
	Token string `json:"token" binding:"required"`
}

type AllResponse struct {
	Data AllData `json:"data"`
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
}

type AllData struct {
	Records []AllRecord `json:"records"`
}

type AllRecord struct {
	Receipt      float64 `json:"receipt"`
	Disbursement float64 `json:"disbursement"`
	Type         string  `json:"type"`
	Time         string  `json:"time"`
}

func (res *AllResponse) setAllResponse(code int, msg string, records []AllRecord) {
	res.Data.Records = records
	res.Code = code
	res.Msg = msg
}

// GetAllBills 获取所有账单
func GetAllBills(db *gorm.DB) func(c *gin.Context) {
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

		u, err := user.GetUserByUsername(db, username)
		if err != nil {
			response.setAllResponse(-3, "获取用户失败，请重新登录", nil)
			c.JSON(http.StatusOK, response)
			return
		}

		billList, err := bill.GetAllBillsByUserID(db, u.ID)
		if err != nil {
			response.setAllResponse(-4, "获取时发生错误，请稍后再试", nil)
			c.JSON(http.StatusOK, response)
		}
		var records []AllRecord
		for _, billRecord := range billList {
			timeRecord := time.Unix(billRecord.Time.Unix(), 0).Format("2006-01-02 15:04:05")
			record := AllRecord{
				Receipt:      billRecord.Receipt,
				Disbursement: billRecord.Disbursement,
				Type:         billRecord.Type,
				Time:         timeRecord,
			}
			records = append(records, record)
		}

		response.setAllResponse(0, "查询成功", records)
		c.JSON(http.StatusOK, response)
	}
}
