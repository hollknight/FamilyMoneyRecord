package modify

import (
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type AdvRequest struct {
	Token              string  `json:"token" binding:"required"`
	AdvanceConsumption float64 `json:"advanceConsumption" binding:"required"`
}

type AdvResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (res *AdvResponse) setAdvResponse(code int, msg string) {
	res.Code = code
	res.Msg = msg
}

// AdvanceConsumption 修改用户预消费金额
func AdvanceConsumption(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(AdvRequest)
		response := new(AdvResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setAdvResponse(-1, "请检查传入参数是否完整")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token := request.Token
		username, err := utils.ParseToken(token)
		if err != nil {
			response.setAdvResponse(-2, "登录已过期，请重新登录")
			c.JSON(http.StatusOK, response)
			return
		}

		advanceConsumption := request.AdvanceConsumption
		err = user.UpdateAdvanceConsumption(db, username, advanceConsumption)
		if err != nil {
			response.setAdvResponse(-3, "预消费金额修改失败")
			c.JSON(http.StatusOK, response)
			return
		}

		response.setAdvResponse(0, "预消费金额修改成功")
		c.JSON(http.StatusOK, response)
	}
}
