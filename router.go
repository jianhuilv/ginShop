package main

import (
	"github.com/gin-gonic/gin"
	"goShop/handler"
	"io"
	"log"
	"os"
)

//初始化gin
func init() {
	//日志
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	log.SetOutput(io.MultiWriter(f, os.Stdout))

	//分配路由
	commonGroup := engine.Group("/common")    //通用使用的接口
	userGroup := engine.Group("/user")        //用户使用的接口
	salesPersonGroup := engine.Group("/sale") //销售员(店家)使用
	adminGroup := engine.Group("/admin")      //  管理者使用的接口

	userGroup.POST("/register", handler.Register)
	userGroup.POST("/login", handler.Login)
	userGroup.GET("/logout", handler.Logout)
	userGroup.GET("/addToCart", handler.Auth, handler.AddToCart)
	userGroup.GET("/pay", handler.Auth, handler.Pay)
	userGroup.GET("/getOrdersByUid", handler.Auth, handler.GetOrdersByUid)

	commonGroup.GET("/getProducts", handler.GetProducts)
	commonGroup.GET("/getProductById", handler.GetProductById)
	commonGroup.GET("/getTypeBySid", handler.GetTypeBySid)

	adminGroup.GET("/reset", handler.ResetPassWordBySid)
	adminGroup.GET("/addSaler", handler.AddSaler)
	adminGroup.GET("/getOrders", handler.GetOrders)
	adminGroup.GET("/getOrdersBySid", handler.GetOrdersBySid)

	salesPersonGroup.GET("/addType", handler.Auth, handler.AddType)
	salesPersonGroup.GET("/deleteType", handler.Auth, handler.DeleteType)
	salesPersonGroup.GET("/insertProduct", handler.Auth, handler.InsertProduct)
	salesPersonGroup.GET("/deleteProduct", handler.Auth, handler.DeleteProduct)
	salesPersonGroup.GET("/updateProduct", handler.Auth, handler.UpdateProduct)
	salesPersonGroup.GET("/sendProduct", handler.Auth, handler.SendProduct)
	salesPersonGroup.GET("/getLogs", handler.Auth, handler.GetLogs)

}
