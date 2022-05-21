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
	//engine.Run()
	//var Db *sql.DB
	//Db, err := sql.Open("mysql", "root:qq69.com@tcp(124.70.87.90:3306)/newshopbase?charset=utf8")
	//if err != nil {
	//	log.Fatal("数据库打开出现了问题：", err)
	//	return
	//}
	//err = db.Ping()
	//if err != nil {
	//	log.Fatal("数据库连接出现了问题：", err)
	//	return
	//}
	//rows, err := db.Query("select * from user")
	//for rows.Next() {
	//
	//	err = rows.Scan(&user1.Uid, &user1.username, &user1.password, &user1.songs)
	//	fmt.Println(user1.uid)
	//	fmt.Println(user1.username)
	//	fmt.Println(user1.password)
	//	fmt.Println(user1.songs)
	//}
	//c, err := redis.Dial("tcp", "127.0.0.1:6379")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//defer c.Close()
	//_, err = c.Do("Set", "key1", 998) //redis写入数据
	//if err != nil {
	//	fmt.Println(err)
	//}
	//r, err := redis.Int(c.Do("Get", "key1")) //类型断言接口
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(r)
	err := engine.Run()
	if err != nil {
		log.Fatalf("cannot run server: %v", err)
	}
}
