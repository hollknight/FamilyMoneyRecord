package user_handlers

import (
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type InfoRequest struct {
	Token string `json:"token" binding:"required"`
}

type InfoResponse struct {
	Data InfoData `json:"data"`
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
}

type InfoData struct {
	Name               string  `json:"name"`
	Username           string  `json:"username"`
	ReceiptSum         float64 `json:"receiptSum"`
	DisbursementSum    float64 `json:"disbursementSum"`
	AdvanceConsumption float64 `json:"AdvanceConsumption"`
}

func (res *InfoResponse) setInfoResponse(code int, name, username, msg string, receiptSum, disbursementSum, AdvanceConsumption float64) {
	res.Data = InfoData{
		Name:               name,
		Username:           username,
		ReceiptSum:         receiptSum,
		DisbursementSum:    disbursementSum,
		AdvanceConsumption: AdvanceConsumption,
	}
	res.Code = code
	res.Msg = msg
}

// GetInfo 获取用户信息接口
func GetInfo(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(InfoRequest)
		response := new(InfoResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setInfoResponse(-1, "", "", "请检查传入参数是否完整", 0, 0, 0)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token := request.Token
		username, err := utils.ParseToken(token)
		if err != nil {
			response.setInfoResponse(-2, "", "", "登录已过期，请重新登录", 0, 0, 0)
			c.JSON(http.StatusOK, response)
			return
		}

		u, err := user.GetUserByUsername(db, username)
		if err != nil {
			response.setInfoResponse(-3, "", "", "未查询到该账号", 0, 0, 0)
			c.JSON(http.StatusOK, response)
			return
		}

		receiptSum, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", u.ReceiptSum), 64)
		disbursementSum, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", u.DisbursementSum), 64)
		advanceConsumption, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", u.AdvanceConsumption), 64)
		response.setInfoResponse(0, u.Name, u.Username, "查询成功", receiptSum, disbursementSum, advanceConsumption)
		c.JSON(http.StatusOK, response)
	}
}
