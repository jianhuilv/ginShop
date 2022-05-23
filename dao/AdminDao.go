package dao

import (
	"goShop/entry"
	"log"
)

func InsertSaler(saler *entry.Saler) (err error) {
	DB := getConn()
	rows, err := DB.Prepare("insert into saler(`sid`,`password`,`sName`) values (?,?,?)")
	if err != nil {
		log.Printf("InsertSaler fail: err=%w", err)
		return
	}
	res, err := rows.Exec(saler.Sid, saler.Password, saler.SName)
	if err != nil {
		log.Printf("InsertSaler fail: res=%v,err=%w", res, err)
		return
	}
	return
}

func UpdatePassword(sid string) (err error) {
	return
}

func SelectOrderBySid(sid string) (orders entry.Order, err error) {
	return
}

func GetOrders() (orders []entry.Order, err error) {
	return
}
