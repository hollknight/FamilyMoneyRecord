package database_handlers

import (
	"FamilyMoneyRecord/config"
	"FamilyMoneyRecord/database/action/account"
	"FamilyMoneyRecord/database/action/bill"
	"FamilyMoneyRecord/database/action/operation"
	"FamilyMoneyRecord/database/action/stock"
	"FamilyMoneyRecord/database/action/user"
	"FamilyMoneyRecord/utils"
	"FamilyMoneyRecord/utils/database_utils"
	"FamilyMoneyRecord/utils/resource_utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type RecoverRequest struct {
	Token    string `json:"token" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type RecoverResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (res *RecoverResponse) setRecoverResponse(code int, msg string) {
	res.Code = code
	res.Msg = msg
}

// RecoverDatabase 清空数据库接口
func RecoverDatabase(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(RecoverRequest)
		response := new(RecoverResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setRecoverResponse(-1, "请检查传入参数是否缺失")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token := request.Token
		password := request.Password
		admin, err := utils.ParseToken(token)
		if err != nil {
			response.setRecoverResponse(-2, "登录已过期，请重新登录")
			c.JSON(http.StatusOK, response)
			return
		}
		if admin != config.AdminUsername {
			response.setRecoverResponse(-3, "权限不足")
			c.JSON(http.StatusOK, response)
			return
		}
		if password != config.AdminPassword {
			response.setRecoverResponse(-4, "密码错误，请重新输入")
			c.JSON(http.StatusOK, response)
			return
		}

		name := request.Name
		path := config.FolderBathURL + name + ".json"
		isExist, err := resource_utils.IsExist(path)
		if isExist || err != nil {
			response.setRecoverResponse(-5, "备份文件不存在，请更换备份文件名称")
			c.JSON(http.StatusOK, response)
			return
		}
		err = db.Transaction(func(tx *gorm.DB) error {
			// 清空数据库表数据
			database_utils.DropAllTables(tx)
			txErr := database_utils.CreateTables(tx)
			if txErr != nil {
				return txErr
			}

			database, txErr := database_utils.JSON2Struct(name)
			if txErr != nil {
				return txErr
			}

			for _, u := range database.Users {
				txErr = user.AddUserByStruct(tx, u)
				if txErr != nil {
					return txErr
				}
			}
			for _, b := range database.Bills {
				txErr = bill.AddBillByStruct(tx, b)
				if txErr != nil {
					return txErr
				}
			}
			for _, a := range database.Accounts {
				txErr = account.AddAccountByStruct(tx, a)
				if txErr != nil {
					return txErr
				}
			}
			for _, s := range database.Stocks {
				txErr = stock.AddStockByStruct(tx, s)
				if txErr != nil {
					return txErr
				}
			}
			for _, o := range database.Operations {
				txErr = operation.AddOperationByStruct(tx, o)
				if txErr != nil {
					return txErr
				}
			}

			response.setRecoverResponse(0, "数据库恢复成功")
			c.JSON(http.StatusOK, response)
			return nil
		})
		if err != nil {
			response.setRecoverResponse(-6, "恢复数据库时发生错误，请稍后再试")
			c.JSON(http.StatusOK, response)
		}
	}
}
