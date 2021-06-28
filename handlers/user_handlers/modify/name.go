package modify

import (
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type NameRequest struct {
	Token string `json:"token" binding:"required"`
	Name  string `json:"name" binding:"required"`
}

type NameResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (res *NameResponse) setNameResponse(code int, msg string) {
	res.Code = code
	res.Msg = msg
}

// Name 修改用户姓名接口
func Name(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(NameRequest)
		response := new(NameResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setNameResponse(-1, "请检查传入参数是否完整")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token := request.Token
		username, err := utils.ParseToken(token)
		if err != nil {
			response.setNameResponse(-2, "登录已过期，请重新登录")
			c.JSON(http.StatusOK, response)
			return
		}

		name := request.Name
		err = user.UpdateUserName(db, username, name)
		if err != nil {
			response.setNameResponse(-3, "姓名修改失败")
			c.JSON(http.StatusOK, response)
			return
		}

		response.setNameResponse(0, "姓名修改成功")
		c.JSON(http.StatusOK, response)
	}
}
