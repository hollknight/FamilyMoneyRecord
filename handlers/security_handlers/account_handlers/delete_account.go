package account_handlers

import (
	"FamilyMoneyRecord/database/action/account"
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type DeleteRequest struct {
	Token    string `json:"token" binding:"required"`
	Password string `json:"password" binding:"required"`
	ID       uint64 `json:"id" binding:"required"`
}

type DeleteResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (res *DeleteResponse) setDeleteResponse(code int, msg string) {
	res.Code = code
	res.Msg = msg
}

// DeleteAccount 删除证券账户
func DeleteAccount(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(DeleteRequest)
		response := new(DeleteResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setDeleteResponse(-1, "请检查传入参数是否缺失")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token := request.Token
		username, err := utils.ParseToken(token)
		if err != nil {
			response.setDeleteResponse(-2, "登录已过期，请重新登录")
			c.JSON(http.StatusOK, response)
			return
		}

		u, err := user.GetUserByUsername(db, username)
		if err != nil {
			response.setDeleteResponse(-3, "获取用户失败，请重新登录")
			c.JSON(http.StatusOK, response)
			return
		}

		password := request.Password
		// 密码验证
		encryptedPassword, err := utils.Encrypt(password)
		if err != nil {
			response.setDeleteResponse(-4, "密码为空，请重新输入")
			c.JSON(http.StatusOK, response)
			return
		}
		if strings.Compare(encryptedPassword, u.Password) != 0 {
			response.setDeleteResponse(-5, "密码错误，请重新输入")
			c.JSON(http.StatusOK, response)
			return
		}

		id := request.ID
		err = account.DeleteAccount(db, id)
		if err != nil {
			response.setDeleteResponse(-6, "删除时发生错误，请稍后再试")
			c.JSON(http.StatusOK, response)
			return
		}

		response.setDeleteResponse(0, "删除成功")
		c.JSON(http.StatusOK, response)
	}
}
