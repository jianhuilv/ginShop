package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"goShop/entry"
	"log"
)

func getConn() (DB *sql.DB) {
	DB, err := sql.Open("mysql", "root:qq69.com@tcp(124.70.87.90:3306)/newshopbase?charset=utf8")
	if err != nil {
		log.Fatal("数据库打开出现了问题：", err)
		return nil
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("数据库连接出现了问题：", err)
		return
	}
	return
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
}
func SelectProductsBySid(pType *entry.PType) (products []entry.Product, err error) {
	DB := getConn()
	rows, err := DB.Query("select * from product where `sid`=? and `typeName`=?", pType.Sid, pType.TypeName)
	if err != nil {
		log.Printf("SelectProductsBySid fail: err=%v", err)
		return
	}

	for rows.Next() {
		var product entry.Product
		err = rows.Scan(&product.Pid, &product.OwnerId, &product.PName, &product.Remain, &product.Price, &product.PType)
		products = append(products, product)
	}
	return
}
func SelectTypeBySid(sid string) (types []entry.PType, err error) {
	DB := getConn()
	rows, err := DB.Query("select * from pType where `sid`=?", sid)
	if err != nil {
		log.Printf("SelectTypeBySid fail: err=%v", err)
		return
	}

	for rows.Next() {
		var pType entry.PType
		err = rows.Scan(&pType.Sid, &pType.TypeName)
		types = append(types, pType)
	}
	return
}
