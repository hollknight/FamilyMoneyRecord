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

type TimeRequest struct {
	Token     string `json:"token" binding:"required"`
	BeginTime string `json:"beginTime" binding:"required"`
	EndTime   string `json:"endTime" binding:"required"`
}

type TimeResponse struct {
	Data TypeData `json:"data"`
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
}

type TimeData struct {
	Records []TypeRecord `json:"records"`
}

type TimeRecord struct {
	Receipt      float64 `json:"receipt"`
	Disbursement float64 `json:"disbursement"`
	Type         string  `json:"type"`
	Time         string  `json:"time"`
}

func (res *TimeResponse) setTimeResponse(code int, msg string, records []TypeRecord) {
	res.Data.Records = records
	res.Code = code
	res.Msg = msg
}

// GetBillsByTime 删除账单
func GetBillsByTime(db *gorm.DB) func(c *gin.Context) {
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

		u, err := user.GetUserByUsername(db, username)
		if err != nil {
			response.setTimeResponse(-3, "获取用户失败，请重新登录", nil)
			c.JSON(http.StatusOK, response)
			return
		}

		beginTimeStr := request.BeginTime
		beginTime, _ := time.Parse("2006-01-02 15:04:05", beginTimeStr)
		endTimeStr := request.EndTime
		endTime, _ := time.Parse("2006-01-02 15:04:05", endTimeStr)
		billList, err := bill.GetAllBills(db, u.ID)
		if err != nil {
			response.setTimeResponse(-4, "获取时发生错误，请稍后再试", nil)
			c.JSON(http.StatusOK, response)
		}
		var records []TypeRecord
		for _, billRecord := range billList {
			if beginTime.Before(billRecord.Time) && endTime.After(billRecord.Time) {
				timeRecord := time.Unix(billRecord.Time.Unix(), 0).Format("2006-01-02 15:04:05")
				record := TypeRecord{
					Receipt:      billRecord.Receipt,
					Disbursement: billRecord.Disbursement,
					Type:         billRecord.Type,
					Time:         timeRecord,
				}
				records = append(records, record)
			}
		}

		response.setTimeResponse(0, "查询成功", records)
		c.JSON(http.StatusOK, response)
	}
}
