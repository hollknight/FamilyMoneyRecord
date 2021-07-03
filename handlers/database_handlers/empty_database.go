package database_handlers

import (
	"FamilyMoneyRecord/config"
	"FamilyMoneyRecord/utils"
	"FamilyMoneyRecord/utils/database_utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type EmptyRequest struct {
	Token    string `json:"token" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type EmptyResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (res *EmptyResponse) setEmptyResponse(code int, msg string) {
	res.Code = code
	res.Msg = msg
}

// EmptyDatabase 清空数据库接口
func EmptyDatabase(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(EmptyRequest)
		response := new(EmptyResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setEmptyResponse(-1, "请检查传入参数是否缺失")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token := request.Token
		password := request.Password
		admin, err := utils.ParseToken(token)
		if err != nil {
			response.setEmptyResponse(-2, "登录已过期，请重新登录")
			c.JSON(http.StatusOK, response)
			return
		}
		if admin != config.AdminUsername {
			response.setEmptyResponse(-3, "权限不足")
			c.JSON(http.StatusOK, response)
			return
		}
		if password != config.AdminPassword {
			response.setEmptyResponse(-4, "密码错误，请重新输入")
			c.JSON(http.StatusOK, response)
			return
		}

		err = db.Transaction(func(tx *gorm.DB) error {
			database_utils.DropAllTables(tx)
			txErr := database_utils.CreateTables(tx)
			if txErr != nil {
				return txErr
			}

			response.setEmptyResponse(0, "数据库清空成功")
			c.JSON(http.StatusOK, response)
			return nil
		})
		if err != nil {
			response.setEmptyResponse(-5, "清空数据库时发生错误，请稍后再试")
			c.JSON(http.StatusOK, response)
		}

	}
}
