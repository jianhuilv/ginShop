package handler

import (
	"github.com/gin-gonic/gin"
	"goShop/dao"
	"goShop/entry"
	"log"
	"net/http"
)

// GetProducts 查询全部商品信息
func GetProducts(c *gin.Context) {
	var pType entry.PType
	err := c.Bind(&pType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
	}
	products, err := dao.SelectProductsBySid(&pType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"flag": "success",
		"data": products,
	})
}

//查询商品信息
func GetProductByPId(c *gin.Context) {
	pid := c.Query("pid")
	product, err := dao.SelectProductByPid(pid)
	if err != nil {
		log.Printf("GetProductByPid")
		c.JSON(500, gin.H{
			"success": false,
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    product,
	})

}

//获取商品类别
func GetTypeBySid(c *gin.Context) {
	sid := c.Query("sid")
	types, err := dao.SelectTypeBySid(sid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"flag": "success",
		"data": types,
	})
}
