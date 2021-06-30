package account_handlers

import (
	"FamilyMoneyRecord/database/action/account"
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type AddRequest struct {
	Token string `json:"token" binding:"required"`
}

type AddResponse struct {
	Data AddData `json:"data"`
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
}

type AddData struct {
	ID uint64 `json:"id"`
}

func (res *AddResponse) setAddResponse(code int, msg string, id uint64) {
	res.Data.ID = id
	res.Code = code
	res.Msg = msg
}

// AddAccount 添加证券账户
func AddAccount(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(AddRequest)
		response := new(AddResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setAddResponse(-1, "请检查传入参数是否缺失", 0)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token := request.Token
		username, err := utils.ParseToken(token)
		if err != nil {
			response.setAddResponse(-2, "登录已过期，请重新登录", 0)
			c.JSON(http.StatusOK, response)
			return
		}

		u, err := user.GetUserByUsername(db, username)
		if err != nil {
			response.setAddResponse(-3, "未查询到该账号", 0)
			c.JSON(http.StatusOK, response)
			return
		}

		id, err := account.AddAccount(db, u.ID)
		if err != nil {
			response.setAddResponse(-4, "添加账户时发生错误", 0)
			c.JSON(http.StatusOK, response)
			return
		}

		response.setAddResponse(0, "添加成功", id)
		c.JSON(http.StatusOK, response)
	}
}
