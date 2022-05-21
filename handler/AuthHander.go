package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
)

// Auth 权限验证
func Auth(ctx *gin.Context) {
	tokenString := ctx.Query("tokenString")
	userId := ctx.Query("uid")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		log.Printf("Auth err: req=%v err=%w", ctx, err)
	}
	if token.Claims.(jwt.MapClaims)["name"] == userId {

	}
}

func secret() jwt.Keyfunc { //按照这样的规则解析
	return func(t *jwt.Token) (interface{}, error) {
		return []byte("gettoken"), nil
	}
}
