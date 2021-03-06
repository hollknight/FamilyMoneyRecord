package user_handlers

import (
	"FamilyMoneyRecord/config"
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
	Token   string `json:"token"`
	IsAdmin bool   `json:"isAdmin"`
}

func (res *LoginResponse) setLoginResponse(code int, token, msg string, isAdmin bool) {
	res.Data = LoginData{
		Token:   token,
		IsAdmin: isAdmin,
	}
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
			response.setLoginResponse(-1, "", "请检查传入参数是否缺失", false)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		username := request.Username
		password := request.Password

		token, err := utils.CreateToken(username)
		if err != nil {
			response.setLoginResponse(-2, "", "生成token时发生错误", false)
			c.JSON(http.StatusOK, response)
			return
		}

		if username == config.AdminUsername && password == config.AdminPassword {
			response.setLoginResponse(0, token, "管理员登录", true)
			c.JSON(http.StatusOK, response)
			return
		}

		u, err := user.GetUserByUsername(db, username)
		if err != nil {
			if username == config.AdminUsername {
				response.setLoginResponse(-3, "", "管理员密码错误", false)
			} else {
				response.setLoginResponse(-4, "", "未查询到该账号", false)
			}
			c.JSON(http.StatusOK, response)
			return
		}

		// 密码验证
		encryptedPassword, err := utils.Encrypt(password)
		if err != nil {
			response.setLoginResponse(-5, "", "密码为空，请重新输入", false)
			c.JSON(http.StatusOK, response)
			return
		}
		if strings.Compare(encryptedPassword, u.Password) != 0 {
			response.setLoginResponse(-6, "", "密码错误，请重新输入", false)
			c.JSON(http.StatusOK, response)
			return
		}

		if u.AdvanceConsumption > 0 && float64(u.DisbursementSum/u.AdvanceConsumption) > 1 {
			response.setLoginResponse(0, token, "登录成功，消费金额超过预消费额", false)
			c.JSON(http.StatusOK, response)
		} else if u.AdvanceConsumption > 0 && float64(u.DisbursementSum/u.AdvanceConsumption) >= 0.8 {
			response.setLoginResponse(0, token, "登录成功，消费金额接近预消费额", false)
			c.JSON(http.StatusOK, response)
		} else {
			response.setLoginResponse(0, token, "登录成功", false)
			c.JSON(http.StatusOK, response)
		}
	}
}
