package dao

import (
	"goShop/entry"
	"log"
)

func InsertProductToDB(product entry.Product) (success bool, err error) {
	DB := getConn()
	rows, err := DB.Prepare("insert into product(`pid`,`ownerId`,`pName`,`detail`,`remain`,`price`,`ptype`) values (?,?,?,?,?,?,?)")
	if err != nil {
		log.Printf("InsertProductToDB fail: err=%v", err)
		return
	}
	res, err := rows.Exec(product.Pid, product.OwnerId, product.PName, product.Detail, product.Remain, product.Price, product.PType)
	if err != nil {
		log.Printf("InsertProductToDB fail: res=%v,err=%v", res, err)
		return
	}
	return
}

func DeleteProductFromDB(pid string) (success bool) {
	DB := getConn()
	stmt, err := DB.Prepare("delete from PType where `pid`=?")
	if err != nil {
		log.Printf("DeleteProductFromDB fail: err=%v", err)
		return
	}
	res, err := stmt.Exec(pid)
	if err != nil {
		log.Printf("DeleteProductFromDB fail: res=%v,err=%v", res, err)
		return
	}
	return
}

func UpdateProductFromDB(product entry.Product) (success bool) {
	DB := getConn()
	statement, err := DB.Prepare("update product SET `pName`=?,`detail`=?,`remain`=?,`price`=?,`ptype`=? where `pid`=?")
	if err != nil {
		log.Printf("UpdateProductFromDB fail: err=%v", err)
		return
	}
	_, err = statement.Exec(product.PName, product.Detail, product.Remain, product.Price, product.PType, product.Pid)
	if err != nil {
		log.Printf("UpdateProductFromDB fail: err=%v", err)
		return
	}
	return
}

func InsertTypeToDB(pType entry.PType) (success bool) {
	DB := getConn()
	rows, err := DB.Prepare("insert into Ptype(`sid`,`typeName`) values (?,?)")
	if err != nil {
		log.Printf("InsertTypeToDB fail: err=%v", err)
		return
	}
	res, err := rows.Exec(pType.Sid, pType.TypeName)
	if err != nil {
		log.Printf("InsertTypeToDB fail: res=%v,err=%v", res, err)
		return
	}
	return
}

func DeleteTypeFromDB(pType entry.PType) (success bool) {
	DB := getConn()
	stmt, err := DB.Prepare("delete from PType where `sid`=?")
	if err != nil {
		log.Printf("DeleteTypeFromDB fail: err=%v", err)
		return
	}
	res, err := stmt.Exec(pType.Sid)
	if err != nil {
		log.Printf("DeleteTypeFromDB fail: res=%v,err=%v", res, err)
		return
	}
	return
}

func GetOptionsFromDB() {

}

func SetOrderSent(order entry.Order) (success bool, err error) {
	DB := getConn()
	_, err = DB.Exec("update order SET `status`=`已发货` where `oid`=?", order.Oid)
	if err != nil {
		log.Printf("UpdatePassword fail: err=%v", err)
		return
	}
	return true, err
}
func SelectSidByUsernameAndPwd(saler entry.Saler) (uuid string) {
	DB := getConn()
	rows, err := DB.Query("select uid from user where username=? and password=?", saler.SName, saler.Password)
	if err != nil {
		log.Printf("SelectUidByUsernameAndPwd fail: err=%w", err)
		return
	}
	for rows.Next() {
		err = rows.Scan(saler.Sid)
		if err != nil {
			log.Printf("SelectUidByUsernameAndPwd fail: err=%w", err)
			return
		}
	}
	return saler.Sid
}
