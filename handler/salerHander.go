package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"goShop/dao"
	"goShop/entry"
	"goShop/utils"
	"log"
	"net/http"
	"time"
)

// LoginForSaler 登陆
func LoginForSaler(c *gin.Context) {
	var saler entry.Saler
	err := c.Bind(&saler)
	if err != nil {
		log.Printf("Login fail: err =%w", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
		})
		return
	}
	saler.Sid = dao.SelectSidByUsernameAndPwd(saler)
	if err != nil {
		log.Printf("Login: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "账号或密码错误",
		})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": saler.Sid,
		"exp":  time.Now().Unix() + 3600*24*3,
		"iss":  saler.Sid,
	})
	tokenString, _ := token.SigningString()
	c.JSON(200, gin.H{
		"flag": "登陆成功",
		"data": map[string]string{
			"uid":   saler.Sid,
			"token": tokenString,
		},
	})
}

// InsertProduct 添加商品
func InsertProduct(c *gin.Context) {
	var product entry.Product
	err1 := c.Bind(&product)
	if err1 != nil {
		log.Printf("InsertProduct fail: err1=%w", err1)
		c.JSON(500, gin.H{
			"success": false,
		})
		return
	}

	// 获取图片
	fileHeader, err := c.FormFile("image")
	if err != nil {
		log.Printf("InsertProduct fail: err1=%w", err1)
		c.JSON(500, gin.H{
			"success": false,
		})
		return
	}
	err = utils.UploadImage(product.Pid, fileHeader)
	if err != nil {
		log.Printf("InsertProduct fail: err1=%w", err1)
		c.JSON(500, gin.H{
			"success": false,
		})
		return
	}

	status, err2 := dao.InsertProductToDB(product)
	if err1 != nil {
		log.Printf("InsertProduct fail: err2=%w", err2)
		c.JSON(500, gin.H{
			"success": status,
		})
		return
	}
	c.JSON(200, gin.H{
		"success": status,
	})
}

// DeleteProduct 删除商品
func DeleteProduct(c *gin.Context) {
	pid := c.Query("pid")
	success := dao.DeleteProductFromDB(pid)
	if !success {
		c.JSON(200, gin.H{
			"success": success,
		})
		return
	}
	c.JSON(200, gin.H{
		"success": success,
	})
}

// UpdateProduct 修改商品信息
func UpdateProduct(c *gin.Context) {
	var product entry.Product
	err := gin.Bind(&product)
	if err != nil {
		log.Printf("UpdateProduct bind fail: err=%w", err)
	}
	success := dao.UpdateProductFromDB(product)

	// 获取图片
	fileHeader, err1 := c.FormFile("image")
	if err1 != nil {
		log.Printf("InsertProduct fail: err1=%v", err1)
		c.JSON(500, gin.H{
			"success": false,
		})
		return
	}
	err2 := utils.UploadImage(product.Pid, fileHeader)
	if err2 != nil {
		log.Printf("InsertProduct fail: err1=%v", err2)
		c.JSON(500, gin.H{
			"success": false,
		})
		return
	}

	if !success {
		c.JSON(200, gin.H{
			"success": success,
		})
		return
	}
	c.JSON(200, gin.H{
		"success": success,
	})
	return

}

// SendProduct 发货
func SendProduct(c *gin.Context) {
	// 接收一个order，返回bool
	var order entry.Order
	err1 := c.Bind(&order)
	status, err2 := dao.SetOrderSent(order)
	if err1 != nil {
		log.Printf("Pay err1=%w err2=%w", err1, err2)
		c.JSON(500, gin.H{})
	}

	c.JSON(200, gin.H{
		"success": status,
	})

}

// AddType 添加商品类别
func AddType(c *gin.Context) {
	var pType entry.PType
	err := c.Bind(pType)
	if err != nil {
		log.Printf("AddType fail err=%w", err)
		c.JSON(200, gin.H{
			"success": false,
		})
	}
	success := dao.InsertTypeToDB(pType)
	if !success {
		c.JSON(200, gin.H{
			"success": success,
		})
	}
}

// DeleteType 删除商品类别
func DeleteType(c *gin.Context) {
	var pType entry.PType
	err := c.Bind(pType)
	if err != nil {
		log.Printf("DeleteType fail err=%w", err)
		c.JSON(200, gin.H{
			"success": false,
		})
	}
	success := dao.DeleteTypeFromDB(pType)
	if !success {
		c.JSON(200, gin.H{
			"success": success,
		})
	}
}

// GetLogs 查询某个商品的所有日志（浏览，购买）
func GetLogs(c *gin.Context) {
	c.Request.Header.Get("X-Forward-For")
}
