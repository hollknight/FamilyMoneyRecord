package user_handlers

import (
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
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
	Name               string `json:"name"`
	Username           string `json:"username"`
	ReceiptSum         int    `json:"receiptSum"`
	DisbursementSum    int    `json:"disbursementSum"`
	AdvanceConsumption int    `json:"AdvanceConsumption"`
}

func (res *InfoResponse) setInfoResponse(code int, name, username, msg string, receiptSum, disbursementSum, AdvanceConsumption int) {
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

		response.setInfoResponse(0, u.Name, u.Username, "查询成功", u.ReceiptSum, u.DisbursementSum, u.AdvanceConsumption)
		c.JSON(http.StatusOK, response)
	}
}
