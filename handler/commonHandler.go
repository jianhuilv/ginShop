package handler

import (
	"github.com/gin-gonic/gin"
	"goShop/dao"
	"goShop/entry"
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
func GetProductById(c *gin.Context) {

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
