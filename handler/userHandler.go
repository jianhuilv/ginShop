package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
	"goShop/dao"
	"goShop/entry"
	"log"
	"net/http"
	"time"
)

var MySecret = []byte("qq.com")

// Register 注册
func Register(c *gin.Context) {

	var user entry.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Printf("Register fail: err =%w", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
		})
		return
	}

	user.Uid = fmt.Sprint(uuid.NewV4())
	err = dao.NewUser(&user)
	if err != nil {
		log.Printf("Register fail: err= %w", err)
		c.JSON(200, gin.H{
			"status":  "fail",
			"Message": "该昵称已被使用",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"status": "success",
			"data":   entry.User{Uid: user.Uid},
		})
	}

}

// Login 登陆
func Login(c *gin.Context) {
	var user entry.User
	err := c.Bind(&user)
	if err != nil {
		log.Printf("Login fail: err =%w", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
		})
		return
	}
	user.Uid = dao.SelectUidByUsernameAndPwd(&user)
	if err != nil {
		log.Printf("Login: %v", err)
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user.Uid,
		"exp":  time.Now().Unix() + 3600*24*3,
		"iss":  user.Uid,
	})
	tokenString, _ := token.SigningString()
	c.JSON(200, gin.H{
		"flag": "登陆成功",
		"data": map[string]string{
			"uid":   user.Uid,
			"token": tokenString,
		},
	})
}

// Logout 注销 交由前端把token删了就可以
func Logout(c *gin.Context) {

}

// AddToCart 添加购物车
func AddToCart(c *gin.Context) {
	var cart entry.Cart
	err1 := c.Bind(&cart)
	err2 := dao.InsertIntoCart(cart)
	if err1 != nil || err2 != nil {
		log.Printf("AddToCart err1=%w, err2=%w", err1, err2)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// Pay 购买
func Pay(c *gin.Context) {
	// 接收一个order，返回bool
	oid := c.Query("oid")
	status, err := dao.UpDateStatusOfOrder(oid)
	if err != nil {
		log.Printf("Pay err1=%v err2=%v", err)
		c.JSON(500, gin.H{})
	}

	c.JSON(200, gin.H{
		"status": status,
	})

}

// GetOrdersByUid 查询自己的订单
func GetOrdersByUid(c *gin.Context) {
	uid := c.Query("uid")
	orders, err := dao.SelectOrdersByUid(uid)
	if err != nil {
		log.Printf("GetOrderByUid: err=%w", err)
		c.JSON(500, gin.H{
			"status": "fail",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
		"data":   orders,
	})
}

// CreatOrder 生成订单
func CreatOrder(c *gin.Context) {
	// 生成一个order
	var order entry.Order
	err := c.Bind(order)
	if err != nil {
		log.Printf("CreatOrder fail: err=%w", err)
		c.JSON(500, gin.H{
			"status": "fail",
		})
		return
	}
	success, err := dao.InsertOrder(order)
	if err != nil {
		log.Printf("CreatOrder fail: err=%w", err)
		c.JSON(500, gin.H{
			"success": false,
		})
		return
	}
	c.JSON(200, gin.H{
		"success": success,
		"orderId": order.Oid,
	})
}
