package modify

import (
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type PasswordRequest struct {
	Token       string `json:"token" binding:"required"`
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

type PasswordResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (res *PasswordResponse) setPasswordResponse(code int, msg string) {
	res.Code = code
	res.Msg = msg
}

// Password 修改用户姓名接口
func Password(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(PasswordRequest)
		response := new(PasswordResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setPasswordResponse(-1, "请检查传入参数是否完整")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token := request.Token
		username, err := utils.ParseToken(token)
		if err != nil {
			response.setPasswordResponse(-2, "登录已过期，请重新登录")
			c.JSON(http.StatusOK, response)
			return
		}

		u, err := user.GetUserByUsername(db, username)
		if err != nil {
			response.setPasswordResponse(-3, "未查询到该账号")
			c.JSON(http.StatusOK, response)
			return
		}

		oldPassword := request.OldPassword
		newPassword := request.NewPassword

		// 密码验证
		encryptedOldPassword, err := utils.Encrypt(oldPassword)
		if err != nil {
			response.setPasswordResponse(-4, "原密码为空，请重新输入")
			c.JSON(http.StatusOK, response)
			return
		}
		if strings.Compare(encryptedOldPassword, u.Password) != 0 {
			response.setPasswordResponse(-5, "密码错误，请重新输入")
			c.JSON(http.StatusOK, response)
			return
		}

		encryptedNewPassword, err := utils.Encrypt(newPassword)
		if err != nil {
			response.setPasswordResponse(-6, "新密码为空，请重新输入")
			c.JSON(http.StatusOK, response)
			return
		}
		err = user.UpdateUserPassword(db, u, encryptedNewPassword)
		if err != nil {
			response.setPasswordResponse(-7, "密码修改失败")
			c.JSON(http.StatusOK, response)
			return
		}

		response.setPasswordResponse(0, "密码修改成功")
		c.JSON(http.StatusOK, response)
	}
}
