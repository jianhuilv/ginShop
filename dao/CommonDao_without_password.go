package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"goShop/entry"
	"goShop/utils"
	"log"
)

func getConnFake() (DB *sql.DB) {
	DB, err := sql.Open("mysql", "password@tcp(ipAddress)/newshopbase?charset=utf8")
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
}
func SelectProductsBySidFake(pType *entry.PType) (products []entry.Product, err error) {
	DB := getConn()
	rows, err := DB.Query("select * from product where `sid`=? and `typeName`=?", pType.Sid, pType.TypeName)
	if err != nil {
		log.Printf("SelectProductsBySid fail: err=%v", err)
		return
	}

	for rows.Next() {
		var product entry.Product
		err = rows.Scan(&product.Pid, &product.OwnerId, &product.PName, &product.Remain, &product.Price, &product.PType)
		product.Image = utils.GetImage(product.Pid)
		products = append(products, product)
	}
	return
}
func SelectTypeBySidFake(sid string) (types []entry.PType, err error) {
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
