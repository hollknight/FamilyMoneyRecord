package database_handlers

import (
	"FamilyMoneyRecord/config"
	"FamilyMoneyRecord/utils"
	"FamilyMoneyRecord/utils/resource_utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteRequest struct {
	Token    string `json:"token" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type DeleteResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (res *DeleteResponse) setDeleteResponse(code int, msg string) {
	res.Code = code
	res.Msg = msg
}

// DeleteDatabase 清空数据库接口
func DeleteDatabase() func(c *gin.Context) {
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
		password := request.Password
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
		if password != config.AdminPassword {
			response.setDeleteResponse(-4, "密码错误，请重新输入")
			c.JSON(http.StatusOK, response)
			return
		}

		name := request.Name
		path := config.FolderBathURL + name + ".json"
		isExist, err := resource_utils.IsExist(path)
		if !isExist || err != nil {
			response.setDeleteResponse(-5, "备份文件已存在，请更换备份文件名称")
			c.JSON(http.StatusOK, response)
			return
		}

		err = resource_utils.DeleteFile(path)
		if err != nil {
			response.setDeleteResponse(-6, "删除数据库备份时发生错误，请刷新后再试")
			c.JSON(http.StatusOK, response)
			return
		}

		response.setDeleteResponse(0, "删除数据库备份成功")
		c.JSON(http.StatusOK, response)
	}
}
