package admin_handlers

import (
	"FamilyMoneyRecord/config"
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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
	Users []Info `json:"users"`
}

type Info struct {
	Username           string  `json:"username"`
	Name               string  `json:"name"`
	ReceiptSum         float64 `json:"receiptSum"`
	DisbursementSum    float64 `json:"disbursementSum"`
	AdvanceConsumption float64 `json:"AdvanceConsumption"`
}

func (res *InfoResponse) setInfoResponse(code int, msg string, users []Info) {
	res.Data.Users = users
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
			response.setInfoResponse(-1, "请检查传入参数是否完整", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token := request.Token
		admin, err := utils.ParseToken(token)
		if err != nil {
			response.setInfoResponse(-2, "登录已过期，请重新登录", nil)
			c.JSON(http.StatusOK, response)
			return
		}
		if admin != config.AdminUsername {
			response.setInfoResponse(-3, "权限不足", nil)
			c.JSON(http.StatusOK, response)
			return
		}

		username := request.Username
		userList, err := user.GetUsersByLikeUsername(db, username)
		if err != nil {
			response.setInfoResponse(-4, "未查询到该账号", nil)
			c.JSON(http.StatusOK, response)
			return
		}

		var usersInfo []Info
		for _, u := range userList {
			receiptSum, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", u.ReceiptSum), 64)
			disbursementSum, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", u.DisbursementSum), 64)
			advanceConsumption, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", u.AdvanceConsumption), 64)

			userInfo := Info{
				Username:           u.Username,
				Name:               u.Name,
				ReceiptSum:         receiptSum,
				DisbursementSum:    disbursementSum,
				AdvanceConsumption: advanceConsumption,
			}
			usersInfo = append(usersInfo, userInfo)
		}

		response.setInfoResponse(0, "查询成功", usersInfo)
		c.JSON(http.StatusOK, response)
	}
}
