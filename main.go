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
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Use(Cors())

	db, err := database.InitDB()
	if err != nil {
		fmt.Println("初始化数据库失败，请检查原因重新启动程序")
		return
	}
	err = database_utils.CreateTables(db)
	if err != nil {
		fmt.Println("数据库动态迁移失败，请检查原因重新启动程序")
		return
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
			adminGroup.POST("/search_info", admin_handlers.GetSingleInfo(db))
			adminGroup.POST("/all_info", admin_handlers.GetAllInfo(db))
			adminGroup.POST("/add", admin_handlers.AddUser(db))
			adminGroup.DELETE("/delete", admin_handlers.DeleteUser(db))
			adminGroup.PUT("/modify_password", admin_handlers.Password(db))
		}
		// 数据库管理路由分组
		databaseGroup := apiGroup.Group("/database")
		{
			databaseGroup.POST("/save", database_handlers.SaveDatabase(db))
			databaseGroup.POST("/recover", database_handlers.RecoverDatabase(db))
			databaseGroup.POST("/get", database_handlers.GetDatabase(db))
			databaseGroup.DELETE("/empty", database_handlers.EmptyDatabase(db))
			databaseGroup.DELETE("/delete", database_handlers.DeleteDatabase())
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

	//监听端口默认为8422
	err = router.Run(":8422")
	if err != nil {
		fmt.Println("初始化路由失败，请检查路由端口是否被占用")
	}
}

// Cors 跨域请求处理中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()
		c.Next()
	}
}
