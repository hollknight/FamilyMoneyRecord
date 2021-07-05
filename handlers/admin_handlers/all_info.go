package admin_handlers

import (
	"FamilyMoneyRecord/config"
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type AllInfoRequest struct {
	Token string `json:"token" binding:"required"`
}

type AllInfoResponse struct {
	Data AllInfoData `json:"data"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
}

type AllInfoData struct {
	Users []UserInfo `json:"users"`
}

type UserInfo struct {
	Username           string  `json:"username"`
	Name               string  `json:"name"`
	ReceiptSum         float64 `json:"receiptSum"`
	DisbursementSum    float64 `json:"disbursementSum"`
	AdvanceConsumption float64 `json:"advanceConsumption"`
}

func (res *AllInfoResponse) setAllInfoResponse(code int, msg string, users []UserInfo) {
	res.Data.Users = users
	res.Code = code
	res.Msg = msg
}

// GetAllInfo 获取用户信息接口
func GetAllInfo(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(AllInfoRequest)
		response := new(AllInfoResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setAllInfoResponse(-1, "请检查传入参数是否完整", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token := request.Token
		admin, err := utils.ParseToken(token)
		if err != nil {
			response.setAllInfoResponse(-2, "登录已过期，请重新登录", nil)
			c.JSON(http.StatusOK, response)
			return
		}
		if admin != config.AdminUsername {
			response.setAllInfoResponse(-3, "权限不足", nil)
		}

		userList, err := user.GetAllUsers(db)
		if err != nil {
			response.setAllInfoResponse(-4, "未查询到该账号", nil)
			c.JSON(http.StatusOK, response)
			return
		}

		var usersInfo []UserInfo
		for _, u := range userList {
			userInfo := UserInfo{
				Username:           u.Username,
				Name:               u.Name,
				ReceiptSum:         u.ReceiptSum,
				DisbursementSum:    u.DisbursementSum,
				AdvanceConsumption: u.AdvanceConsumption,
			}
			usersInfo = append(usersInfo, userInfo)
		}

		response.setAllInfoResponse(0, "查询成功", usersInfo)
		c.JSON(http.StatusOK, response)
	}
}
