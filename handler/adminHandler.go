package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"goShop/dao"
	"goShop/entry"
)

//添加销售人员
func AddSaler(c *gin.Context) {
	var saler entry.Saler
	saler.SName = c.Query("sName")
	saler.Sid = fmt.Sprint(uuid.NewV4())
	err := dao.InsertSaler(&saler)
	if err != nil {
		c.JSON(200, gin.H{
			"flag": "fail",
			"Data": "该名已被使用",
		})
		return
	}
	c.JSON(200, gin.H{
		"flag": "success",
		"Data": saler,
	})
}

// todo 重置销售人员密码
func ResetPassWordBySid(c *gin.Context) {

}

//查询某个销售人员的各个类别的销售情况
func GetOrdersBySid(c *gin.Context) {

}

//获取所有订单信息
func GetOrders(c *gin.Context) {

}
