package dao

import (
	"goShop/entry"
	"log"
)

func InsertSaler(saler *entry.Saler) (err error) {
	DB := getConn()
	rows, err := DB.Prepare("insert into saler(`sid`,`password`,`sName`) values (?,?,?)")
	if err != nil {
		log.Printf("InsertSaler fail: err=%v", err)
		return
	}
	res, err := rows.Exec(saler.Sid, saler.Password, saler.SName)
	if err != nil {
		log.Printf("InsertSaler fail: res=%v,err=%v", res, err)
		return
	}
	return
}

func UpdatePassword(sid string) (err error) {
	DB := getConn()
	_, err = DB.Exec("update saler SET `password`=`12356` where `sid`=?", sid)
	if err != nil {
		log.Printf("UpdatePassword fail: err=%v", err)
		return
	}

	return
}

func SelectOrdersBySid(sid string) (orders []entry.Order, err error) {
	DB := getConn()
	rows, err := DB.Query("select * from orders where `sid`=?", sid)
	if err != nil {
		log.Printf("UpdatePassword fail: err=%v", err)
		return
	}

	for rows.Next() {
		var order entry.Order
		err = rows.Scan(&order.Oid, &order.Pid, &order.Time, &order.Amount)
		orders = append(orders, order)
	}
	return
}

func GetOrders() (orders []entry.Order, err error) {
	DB := getConn()
	rows, err := DB.Query("select * from orders")
	if err != nil {
		log.Printf("UpdatePassword fail: err=%v", err)
		return
	}

	for rows.Next() {
		var order entry.Order
		err = rows.Scan(&order.Oid, &order.Pid, &order.Time, &order.Amount)
		orders = append(orders, order)
	}
	return
}
