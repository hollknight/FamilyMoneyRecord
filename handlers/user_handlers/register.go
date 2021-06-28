package user_handlers

import (
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type RegisterResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (res *RegisterResponse) setRegisterResponse(code int, msg string) {
	res.Code = code
	res.Msg = msg
}

// Register 用户注册接口
func Register(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(RegisterRequest)
		response := new(RegisterResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setRegisterResponse(-1, "请检查传入参数是否缺失")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		username := request.Username
		password := request.Password
		name := request.Name

		encryptedPassword, err := utils.Encrypt(password)
		if err != nil {
			response.setRegisterResponse(-4, "密码不能空，请重新输入")
			c.JSON(http.StatusOK, response)
			return
		}
		err = user.AddUser(db, username, encryptedPassword, name)
		if err != nil {
			response.setRegisterResponse(-5, "注册失败，请重新注册")
			c.JSON(http.StatusOK, response)
			return
		}

		response.setRegisterResponse(0, "注册成功")
		c.JSON(http.StatusOK, response)
	}
}
