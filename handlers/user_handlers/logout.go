package user_handlers

import (
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type LogoutRequest struct {
	Token    string `json:"token" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LogoutResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (res *LogoutResponse) setLogoutResponse(code int, msg string) {
	res.Code = code
	res.Msg = msg
}

func Logout(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(LogoutRequest)
		response := new(LogoutResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setLogoutResponse(-1, "请检查传入参数是否缺失")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token := request.Token
		password := request.Password
		username, err := utils.ParseToken(token)
		if err != nil {
			response.setLogoutResponse(-2, "登录已过期，请重新登录")
			c.JSON(http.StatusOK, response)
			return
		}

		u, err := user.GetUserByUsername(db, username)
		if err != nil {
			response.setLogoutResponse(-3, "获取用户失败，请重新登录")
			c.JSON(http.StatusOK, response)
			return
		}

		encryptedPassword, _ := utils.Encrypt(password)
		if strings.Compare(encryptedPassword, u.Password) != 0 {
			response.setLogoutResponse(-4, "密码错误，请重新输入")
			c.JSON(http.StatusOK, response)
			return
		}

		err = user.DeleteUser(db, u)
		if err != nil {
			response.setLogoutResponse(-5, "注销时发生错误，请稍后再试")
			c.JSON(http.StatusOK, response)
			return
		}

		response.setLogoutResponse(0, "注销成功")
		c.JSON(http.StatusOK, response)
	}
}
