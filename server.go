package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var engine = gin.Default()

func main() {
	err := engine.Run(":8083")
	if err != nil {
		log.Fatalf("cannot run server: %v", err)
	}
}
