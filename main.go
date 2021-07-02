package main

import (
	"FamilyMoneyRecord/database"
	"FamilyMoneyRecord/handlers/admin_handlers"
	"FamilyMoneyRecord/handlers/bill_handlers"
	"FamilyMoneyRecord/handlers/database_handlers"
	"FamilyMoneyRecord/handlers/security_handlers/account_handlers"
	"FamilyMoneyRecord/handlers/security_handlers/operation_handlers"
	"FamilyMoneyRecord/handlers/security_handlers/stock_handlers"
	"FamilyMoneyRecord/handlers/user_handlers"
	"FamilyMoneyRecord/handlers/user_handlers/modify"
	"FamilyMoneyRecord/utils/database_utils"
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
	err = database_utils.CreateTables(db)
	if err != nil {
		fmt.Println("数据库动态迁移失败，请检查原因重新启动程序")
	}

	// 测试服务是否成功启动路由
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	})

	apiGroup := router.Group("/api")
	{
		// 用户管理路由分组
		userGroup := apiGroup.Group("/user")
		{
			userGroup.POST("/register", user_handlers.Register(db))
			userGroup.POST("/login", user_handlers.Login(db))
			userGroup.POST("/info", user_handlers.GetInfo(db))
			userGroup.DELETE("/logout", user_handlers.Logout(db))
			// 修改用户信息相关路由分组
			modifyGroup := userGroup.Group("/modify")
			{
				modifyGroup.PUT("/name", modify.Name(db))
				modifyGroup.PUT("/password", modify.Password(db))
				modifyGroup.PUT("/advance_consumption", modify.AdvanceConsumption(db))
			}
		}
		// 管理员用户管理路由分组
		adminGroup := apiGroup.Group("/admin")
		{
			adminGroup.POST("/single_info", admin_handlers.GetSingleInfo(db))
			adminGroup.POST("/all_info", admin_handlers.GetAllInfo(db))
			adminGroup.POST("/add", admin_handlers.AddUser(db))
			adminGroup.DELETE("/delete", admin_handlers.DeleteUser(db))
			adminGroup.PUT("/modify_password", admin_handlers.Password(db))
		}
		// 数据库管理路由分组
		databaseGroup := apiGroup.Group("/database")
		{
			databaseGroup.DELETE("/empty", database_handlers.EmptyDatabase(db))
		}
		// 用户收支管理路由分组
		billGroup := apiGroup.Group("/bill")
		{
			billGroup.POST("/add", bill_handlers.AddBill(db))
			billGroup.POST("/get_by_type", bill_handlers.GetBillsByType(db))
			billGroup.POST("/get_by_time", bill_handlers.GetBillsByTime(db))
			billGroup.POST("/get_all", bill_handlers.GetAllBills(db))
			billGroup.DELETE("/delete", bill_handlers.DeleteBill(db))
			billGroup.PUT("/update", bill_handlers.UpdateBill(db))
		}
		// 财务管理路由分组
		securityGroup := apiGroup.Group("/security")
		{
			// 证券账户路由管理分组
			accountGroup := securityGroup.Group("/account")
			{
				accountGroup.POST("/add", account_handlers.AddAccount(db))
				accountGroup.POST("/get", account_handlers.GetAllAccounts(db))
				accountGroup.DELETE("/delete", account_handlers.DeleteAccount(db))
			}
			// 股票路由分组
			stockGroup := securityGroup.Group("/stock")
			{
				stockGroup.POST("/get", stock_handlers.GetAllStocks(db))
			}
			// 股票操作记录路由分组
			operationGroup := securityGroup.Group("/operation")
			{
				operationGroup.POST("/add", operation_handlers.AddOperation(db))
				operationGroup.POST("/get_all", operation_handlers.GetAllOperations(db))
				operationGroup.POST("/get_by_time", operation_handlers.GetOperationsByTime(db))
				operationGroup.DELETE("/delete", operation_handlers.DeleteOperation(db))
				operationGroup.PUT("/update", operation_handlers.UpdateOperation(db))
			}
		}
	}

	//监听端口默认为8421
	err = router.Run(":8422")
	if err != nil {
		fmt.Println("初始化路由失败，请检查路由端口是否被占用")
	}
}
