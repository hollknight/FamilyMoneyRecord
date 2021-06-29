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

type TypeRequest struct {
	Token string `json:"token" binding:"required"`
	Type  string `json:"type" binding:"required"`
}

type TypeResponse struct {
	Data TypeData `json:"data"`
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
}

type TypeData struct {
	Records []TypeRecord `json:"records"`
}

type TypeRecord struct {
	Receipt      int    `json:"receipt"`
	Disbursement int    `json:"disbursement"`
	Type         string `json:"type"`
	Time         string `json:"time"`
}

func (res *TypeResponse) setTypeResponse(code int, msg string, records []TypeRecord) {
	res.Data.Records = records
	res.Code = code
	res.Msg = msg
}

// GetBillsByType 删除账单
func GetBillsByType(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(TypeRequest)
		response := new(TypeResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setTypeResponse(-1, "请检查传入参数是否缺失", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token := request.Token
		username, err := utils.ParseToken(token)
		if err != nil {
			response.setTypeResponse(-2, "登录已过期，请重新登录", nil)
			c.JSON(http.StatusOK, response)
			return
		}

		u, err := user.GetUserByUsername(db, username)
		if err != nil {
			response.setTypeResponse(-3, "获取用户失败，请重新登录", nil)
			c.JSON(http.StatusOK, response)
			return
		}

		moneyType := request.Type
		billList, err := bill.GetBillsByType(db, u.ID, moneyType)
		if err != nil {
			response.setTypeResponse(-4, "获取时发生错误，请稍后再试", nil)
			c.JSON(http.StatusOK, response)
		}
		var records []TypeRecord
		for _, billRecord := range billList {
			timeRecord := time.Unix(billRecord.Time.Unix(), 0).Format("2006-01-02 15:04:05")
			billList := TypeRecord{
				Receipt:      billRecord.Receipt,
				Disbursement: billRecord.Disbursement,
				Type:         billRecord.Type,
				Time:         timeRecord,
			}
			records = append(records, billList)
		}

		response.setTypeResponse(0, "查询成功", records)
		c.JSON(http.StatusOK, response)
	}
}
