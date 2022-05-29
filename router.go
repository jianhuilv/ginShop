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
	userGroup.GET("/addToCart", handler.AuthUser, handler.AddToCart)
	userGroup.GET("/pay", handler.AuthUser, handler.Pay)
	userGroup.GET("/newOrder", handler.AuthUser, handler.CreatOrder)
	userGroup.GET("/getOrdersByUid", handler.AuthUser, handler.GetOrdersByUid)

	commonGroup.GET("/getProducts", handler.GetProducts)
	commonGroup.GET("/getProductById", handler.GetProductByPId)
	commonGroup.GET("/getTypeBySid", handler.GetTypeBySid)

	adminGroup.GET("/reset", handler.ResetPassWordBySid)
	adminGroup.GET("/addSaler", handler.AddSaler)
	adminGroup.GET("/getOrders", handler.GetOrders)
	adminGroup.GET("/getOrdersBySid", handler.GetOrdersBySid)

	salesPersonGroup.GET("/addType", handler.AuthSaler, handler.AddType)
	salesPersonGroup.GET("/deleteType", handler.AuthSaler, handler.DeleteType)
	salesPersonGroup.GET("/insertProduct", handler.AuthSaler, handler.InsertProduct)
	salesPersonGroup.GET("/deleteProduct", handler.AuthSaler, handler.DeleteProduct)
	salesPersonGroup.GET("/updateProduct", handler.AuthSaler, handler.UpdateProduct)
	salesPersonGroup.GET("/sendProduct", handler.AuthSaler, handler.SendProduct)
	salesPersonGroup.GET("/getLogs", handler.AuthSaler, handler.GetLogs)

}
