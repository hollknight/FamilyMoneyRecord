package admin_handlers

import (
	"FamilyMoneyRecord/config"
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type PasswordRequest struct {
	Token    string `json:"token" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PasswordResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (res *PasswordResponse) setPasswordResponse(code int, msg string) {
	res.Code = code
	res.Msg = msg
}

// Password 修改用户密码接口
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
		admin, err := utils.ParseToken(token)
		if err != nil {
			response.setPasswordResponse(-2, "登录已过期，请重新登录")
			c.JSON(http.StatusOK, response)
			return
		}
		if admin != config.AdminUsername {
			response.setPasswordResponse(-3, "权限不足")
			c.JSON(http.StatusOK, response)
			return
		}

		username := request.Username
		u, err := user.GetUserByUsername(db, username)
		if err != nil {
			response.setPasswordResponse(-4, "未查询到该账号")
			c.JSON(http.StatusOK, response)
			return
		}

		password := request.Password
		encryptedPassword, err := utils.Encrypt(password)
		if err != nil {
			response.setPasswordResponse(-5, "新密码为空，请重新输入")
			c.JSON(http.StatusOK, response)
			return
		}
		err = user.UpdateUserPassword(db, u, encryptedPassword)
		if err != nil {
			response.setPasswordResponse(-6, "密码修改失败")
			c.JSON(http.StatusOK, response)
			return
		}

		response.setPasswordResponse(0, "密码修改成功")
		c.JSON(http.StatusOK, response)
	}
}
