package user_handlers

import (
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Data LoginData `json:"data"`
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
}

type LoginData struct {
	Token string `json:"token"`
}

func (res *LoginResponse) setLoginResponse(code int, token, msg string) {
	res.Data.Token = token
	res.Code = code
	res.Msg = msg
}

// Login 用户登录接口
func Login(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(LoginRequest)
		response := new(LoginResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setLoginResponse(-1, "", "请检查传入参数是否缺失")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		username := request.Username
		password := request.Password

		u, err := user.GetUserByUsername(db, username)
		if err != nil {
			response.setLoginResponse(-2, "", "未查询到该账号")
			c.JSON(http.StatusOK, response)
			return
		}

		// 密码验证
		encryptedPassword, err := utils.Encrypt(password)
		if err != nil {
			response.setLoginResponse(-3, "", "密码为空，请重新输入")
			c.JSON(http.StatusOK, response)
			return
		}
		if strings.Compare(encryptedPassword, u.Password) != 0 {
			response.setLoginResponse(-4, "", "密码错误，请重新输入")
			c.JSON(http.StatusOK, response)
			return
		}

		token, err := utils.CreateToken(username)
		if err != nil {
			response.setLoginResponse(-5, "", "生成token时发生错误")
			c.JSON(http.StatusOK, response)
			return
		}

		response.setLoginResponse(0, token, "登录成功")
		c.JSON(http.StatusOK, response)
	}
}
