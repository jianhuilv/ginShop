package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"goShop/entry"
	"log"
)

var user1 entry.User
var engine = gin.Default()
var Db sql.DB

func main() {
	err := engine.Run()
	if err != nil {
		log.Fatalf("cannot run server: %v", err)
	}
}
