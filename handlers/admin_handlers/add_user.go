package admin_handlers

import (
	"FamilyMoneyRecord/config"
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type AddUserRequest struct {
	Token    string `json:"token" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type AddUserResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (res *AddUserResponse) setAddUserResponse(code int, msg string) {
	res.Code = code
	res.Msg = msg
}

// AddUser 用户注册接口
func AddUser(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(AddUserRequest)
		response := new(AddUserResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setAddUserResponse(-1, "请检查传入参数是否缺失")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token := request.Token
		admin, err := utils.ParseToken(token)
		if err != nil {
			response.setAddUserResponse(-2, "登录已过期，请重新登录")
			c.JSON(http.StatusOK, response)
			return
		}
		if admin != config.AdminUsername {
			response.setAddUserResponse(-3, "权限不足")
			c.JSON(http.StatusOK, response)
			return
		}

		username := request.Username
		password := request.Password
		name := request.Name

		encryptedPassword, err := utils.Encrypt(password)
		if err != nil {
			response.setAddUserResponse(-4, "密码不能空，请重新输入")
			c.JSON(http.StatusOK, response)
			return
		}
		err = user.AddUser(db, username, encryptedPassword, name)
		if err != nil {
			response.setAddUserResponse(-5, "注册失败，请重新注册")
			c.JSON(http.StatusOK, response)
			return
		}

		response.setAddUserResponse(0, "注册成功")
		c.JSON(http.StatusOK, response)
	}
}
