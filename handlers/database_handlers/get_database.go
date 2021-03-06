package database_handlers

import (
	"FamilyMoneyRecord/config"
	"FamilyMoneyRecord/utils"
	"FamilyMoneyRecord/utils/resource"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type GetRequest struct {
	Token string `json:"token" binding:"required"`
}

type GetResponse struct {
	GetData GetData `json:"data"`
	Code    int     `json:"code"`
	Msg     string  `json:"msg"`
}

type GetData struct {
	Save []Name `json:"save"`
}

type Name struct {
	Name string `json:"name"`
}

func (res *GetResponse) setGetResponse(code int, msg string, names []Name) {
	res.GetData.Save = names
	res.Code = code
	res.Msg = msg
}

// GetDatabase 清空数据库接口
func GetDatabase(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		request := new(GetRequest)
		response := new(GetResponse)

		err := c.BindJSON(&request)
		if err != nil {
			response.setGetResponse(-1, "请检查传入参数是否缺失", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		token := request.Token
		admin, err := utils.ParseToken(token)
		if err != nil {
			response.setGetResponse(-2, "登录已过期，请重新登录", nil)
			c.JSON(http.StatusOK, response)
			return
		}
		if admin != config.AdminUsername {
			response.setGetResponse(-3, "权限不足", nil)
			c.JSON(http.StatusOK, response)
			return
		}

		var nameInfo []Name
		names, err := resource.GetJSONName()
		for _, name := range names {
			n := Name{
				Name: name,
			}
			nameInfo = append(nameInfo, n)
		}
		if err != nil {
			response.setGetResponse(-4, "获取失败，请稍后再试", nil)
			c.JSON(http.StatusOK, response)
			return
		}

		response.setGetResponse(0, "获取成功", nameInfo)
		c.JSON(http.StatusOK, response)
	}
}
