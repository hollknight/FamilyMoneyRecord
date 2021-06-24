package main

import (
	"FamilyMoneyRecord/database"
	bill2 "FamilyMoneyRecord/database/models/bill"
	user2 "FamilyMoneyRecord/database/models/user"
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
	err = db.AutoMigrate(&user2.User{}, &bill2.Bill{})
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
	router.Run(":8422")
}
