package dao

import (
	"goShop/entry"
	"log"
)

func InsertProductToDB(product entry.Product) (success bool, err error) {
	DB := getConn()
	rows, err := DB.Prepare("insert into product(`pid`,`ownerId`,`pName`,`detail`,`remain`,`price`,`ptype`) values (?,?,?,?,?,?,?)")
	if err != nil {
		log.Printf("InsertProductToDB fail: err=%w", err)
		return
	}
	res, err := rows.Exec(product.Pid, product.OwnerId, product.PName, product.Detail, product.Remain, product.Price, product.PType)
	if err != nil {
		log.Printf("InsertProductToDB fail: res=%v,err=%w", res, err)
		return
	}
	return
}

func DeleteProductFromDB(pid string) (success bool) {
	return
}

func UpdateProductFromDB(product entry.Product) (success bool) {
	return
}

func InsertTypeToDB(pType entry.PType) (success bool) {
	DB := getConn()
	rows, err := DB.Prepare("insert into Ptype(`sid`,`typeName`) values (?,?)")
	if err != nil {
		log.Printf("InsertTypeToDB fail: err=%w", err)
		return
	}
	res, err := rows.Exec(pType.Sid, pType.TypeName)
	if err != nil {
		log.Printf("InsertTypeToDB fail: res=%v,err=%w", res, err)
		return
	}
	return
}

func DeleteTypeFromDB(pType entry.PType) (success bool) {
	return
}

func GetOptionsFromDB() {

}

func SetOrderSent(order entry.Order) (success bool, err error) {
	return
}
func SelectSidByUsernameAndPwd(saler entry.Saler) (uuid string) {
	return
}
