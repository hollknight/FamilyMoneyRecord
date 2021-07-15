package database_handlers

import (
	"FamilyMoneyRecord/config"
	"FamilyMoneyRecord/utils"
	"FamilyMoneyRecord/utils/database"
	"FamilyMoneyRecord/utils/resource"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type SaveRequest struct {
	Token string `json:"token" binding:"required"`
	Name  string `json:"name" binding:"required"`
}

type SaveResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (res *SaveResponse) setSaveResponse(code int, msg string) {
	res.Code = code
	res.Msg = msg
}

// SaveDatabase 清空数据库接口
func SaveDatabase(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(SaveRequest)
		response := new(SaveResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setSaveResponse(-1, "请检查传入参数是否缺失")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token := request.Token
		admin, err := utils.ParseToken(token)
		if err != nil {
			response.setSaveResponse(-2, "登录已过期，请重新登录")
			c.JSON(http.StatusOK, response)
			return
		}
		if admin != config.AdminUsername {
			response.setSaveResponse(-3, "权限不足")
			c.JSON(http.StatusOK, response)
			return
		}

		name := request.Name
		path := config.FolderBathURL + name + ".json"
		isExist, err := resource.IsExist(path)
		if isExist || err != nil {
			response.setSaveResponse(-4, "备份文件已存在，请更换备份文件名称")
			c.JSON(http.StatusOK, response)
			return
		}

		saveDatabase, err := database.SaveDatabase(db)
		if err != nil {
			response.setSaveResponse(-5, "备份失败，请重新再试")
			c.JSON(http.StatusOK, response)
			return
		}
		err = database.Struct2JSON(saveDatabase, name)
		if err != nil {
			response.setSaveResponse(-6, "备份失败，请重新再试")
			c.JSON(http.StatusOK, response)
			return
		}

		response.setSaveResponse(0, "备份成功")
		c.JSON(http.StatusOK, response)
	}
}
