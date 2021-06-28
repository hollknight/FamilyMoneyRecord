package admin_handlers

import (
	"FamilyMoneyRecord/config"
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type DeleteRequest struct {
	Token    string `json:"token" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type DeleteResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (res *DeleteResponse) setDeleteResponse(code int, msg string) {
	res.Code = code
	res.Msg = msg
}

// DeleteUser 删除用户
func DeleteUser(db *gorm.DB) func(c *gin.Context) {
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
		admin, err := utils.ParseToken(token)
		if err != nil {
			response.setDeleteResponse(-2, "登录已过期，请重新登录")
			c.JSON(http.StatusOK, response)
			return
		}
		if admin != config.AdminUsername {
			response.setDeleteResponse(-3, "权限不足")
			c.JSON(http.StatusOK, response)
			return
		}

		username := request.Username
		u, err := user.GetUserByUsername(db, username)
		if err != nil {
			response.setDeleteResponse(-4, "未查询到该账号")
			c.JSON(http.StatusOK, response)
			return
		}

		err = user.DeleteUser(db, u)
		if err != nil {
			response.setDeleteResponse(-5, "删除用户时发生错误，请稍后再试")
			c.JSON(http.StatusOK, response)
			return
		}

		response.setDeleteResponse(0, "删除成功")
		c.JSON(http.StatusOK, response)
	}
}
