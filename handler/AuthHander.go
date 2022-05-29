package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
)

// AuthUser 权限验证
func AuthUser(ctx *gin.Context) {
	tokenString := ctx.Query("tokenString")
	userId := ctx.Query("uid")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		log.Printf("AuthUser err: req=%v err=%v", ctx, err)
		ctx.JSON(200, gin.H{
			"status": "fail",
			"msg":    "token wrong or have expired",
		})
	}
	if token.Claims.(jwt.MapClaims)["name"] != userId {
		log.Printf("AuthUser err: req=%v err=%v", ctx, err)
		ctx.JSON(200, gin.H{
			"status": "fail",
			"msg":    "you have no right to do this",
		})
	}
}

func AuthSaler(ctx *gin.Context) {
	tokenString := ctx.Query("tokenString")
	userId := ctx.Query("uid")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		log.Printf("AuthUser err: req=%v err=%v", ctx, err)
		ctx.JSON(200, gin.H{
			"status": "fail",
			"msg":    "token wrong or have expired",
		})
	}
	if token.Claims.(jwt.MapClaims)["name"] != userId {
		log.Printf("AuthUser err: req=%v err=%v", ctx, err)
		ctx.JSON(200, gin.H{
			"status": "fail",
			"msg":    "you have no right to do this",
		})
	}
}
