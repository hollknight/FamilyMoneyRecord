package bill_handlers

import (
	"FamilyMoneyRecord/database/action/bill"
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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
	ID           uint64  `json:"id"`
	Receipt      float64 `json:"receipt"`
	Disbursement float64 `json:"disbursement"`
	Type         string  `json:"type"`
	Time         string  `json:"time"`
	Note         string  `json:"note"`
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
			receipt, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", billRecord.Receipt), 64)
			disbursement, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", billRecord.Disbursement), 64)
			timeRecord := time.Unix(billRecord.Time.Unix(), 0).Format("2006-01-02 15:04:05")
			record := TypeRecord{
				ID:           billRecord.ID,
				Receipt:      receipt,
				Disbursement: disbursement,
				Type:         billRecord.Type,
				Time:         timeRecord,
				Note:         billRecord.Note,
			}
			records = append(records, record)
		}

		response.setTypeResponse(0, "查询成功", records)
		c.JSON(http.StatusOK, response)
	}
}
