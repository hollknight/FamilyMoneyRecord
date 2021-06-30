package admin_handlers

import (
	"FamilyMoneyRecord/config"
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type InfoRequest struct {
	Token    string `json:"token" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type InfoResponse struct {
	Data InfoData `json:"data"`
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
}

type InfoData struct {
	Name               string  `json:"name"`
	ReceiptSum         float64 `json:"receiptSum"`
	DisbursementSum    float64 `json:"disbursementSum"`
	AdvanceConsumption float64 `json:"AdvanceConsumption"`
}

func (res *InfoResponse) setInfoResponse(code int, name, msg string, receiptSum, disbursementSum, AdvanceConsumption float64) {
	res.Data = InfoData{
		Name:               name,
		ReceiptSum:         receiptSum,
		DisbursementSum:    disbursementSum,
		AdvanceConsumption: AdvanceConsumption,
	}
	res.Code = code
	res.Msg = msg
}

// GetSingleInfo 获取用户信息接口
func GetSingleInfo(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(InfoRequest)
		response := new(InfoResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setInfoResponse(-1, "", "请检查传入参数是否完整", 0, 0, 0)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token := request.Token
		admin, err := utils.ParseToken(token)
		if err != nil {
			response.setInfoResponse(-2, "", "登录已过期，请重新登录", 0, 0, 0)
			c.JSON(http.StatusOK, response)
			return
		}
		if admin != config.AdminUsername {
			response.setInfoResponse(-3, "", "权限不足", 0, 0, 0)
			c.JSON(http.StatusOK, response)
			return
		}

		username := request.Username
		u, err := user.GetUserByUsername(db, username)
		if err != nil {
			response.setInfoResponse(-4, "", "未查询到该账号", 0, 0, 0)
			c.JSON(http.StatusOK, response)
			return
		}

		response.setInfoResponse(0, u.Name, "查询成功", u.ReceiptSum, u.DisbursementSum, u.AdvanceConsumption)
		c.JSON(http.StatusOK, response)
	}
}
