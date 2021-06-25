package main

import (
	"FamilyMoneyRecord/database"
	"FamilyMoneyRecord/database/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	db, err := database.InitDB()
	if err != nil {
		fmt.Println("初始化数据库失败，请检查原因重新启动程序")
		return
	}
	err = db.AutoMigrate(&model.User{}, &model.Bill{}, &model.Account{}, &model.Operation{}, &model.Stock{})
	if err != nil {
		fmt.Println("数据库动态迁移失败，请检查原因重新启动程序")
	}

	// 测试服务是否成功启动路由
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	})

	//监听端口默认为8421
	err = router.Run(":8422")
	if err != nil {
		fmt.Println("初始化路由失败，请检查路由端口是否被占用")
	}
}
