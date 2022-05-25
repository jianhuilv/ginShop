package dao

import (
	"goShop/entry"
	"log"
)

func NewUser(user *entry.User) (err error) {
	DB := getConn()
	rows, err := DB.Prepare("insert into user(`uid`,`mail`,`username`,`password`) values (?,?,?,?)")
	if err != nil {
		log.Printf("execute Insert new user fail: err=%w", err)
		return
	}
	res, err := rows.Exec(user.Uid, user.Mail, user.Username, user.Password)
	if err != nil {
		log.Printf("Insert new user fail: res=%v,err=%w", res, err)
		return
	}
	return
}

func SelectUidByUsernameAndPwd(user *entry.User) (uuid string) {
	DB := getConn()
	rows, err := DB.Query("select uid from user where username=? and password=?", user.Username, user.Password)
	if err != nil {
		log.Printf("SelectUidByUsernameAndPwd fail: err=%w", err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&user.Uid)
		if err != nil {
			log.Printf("SelectUidByUsernameAndPwd fail: err=%w", err)
			return
		}
	}
	return user.Uid
}

func InsertIntoCart(cart entry.Cart) (err error) {
	DB := getConn()
	rows, err := DB.Prepare("insert into cart(`pid`,`ownerId`,`amount`) values (?,?,?)")
	if err != nil {
		log.Printf("InsertIntoCart fail: err=%w", err)
		return
	}
	res, err := rows.Exec(cart.Pid, cart.OwnerId, cart.Amount)
	if err != nil {
		log.Printf("InsertIntoCart fail: res=%v,err=%v", res, err)
		return
	}
	return
}

func InsertOrder(order entry.Order) (success bool, err error) {
	DB := getConn()
	rows, err := DB.Prepare("insert into orders(`oid`,`pid`,`amount`,`status`) values (?,?,?,?,?)")
	if err != nil {
		log.Printf("InsertIntoCart fail: err=%v", err)
		return
	}
	res, err := rows.Exec(order.Oid, order.Pid, order.Amount, "尚未付款")
	if err != nil {
		log.Printf("InsertIntoCart fail: res=%v,err=%v", res, err)
		return
	}
	return
}

func UpDateStatusOfOrder(oid string) (success bool, err error) {
	DB := getConn()
	_, err = DB.Exec("update order SET `status`=`已付款` where `oid`=?", oid)
	if err != nil {
		log.Printf("UpdatePassword fail: err=%v", err)
		return
	}
	return
}

func SelectOrdersByUid(uid string) (orders []entry.Order, err error) {
	DB := getConn()
	rows, err := DB.Query("select * from orders where `uid`=?", uid)
	if err != nil {
		log.Printf("SelectOrdersByUid fail: err=%v", err)
		return
	}

	for rows.Next() {
		var order entry.Order
		err = rows.Scan(&order.Oid, &order.Pid, &order.Amount, &order.Time, &order.Uid)
		orders = append(orders, order)
	}
	return
}

func SelectProductByPid(pid string) (product entry.Product, err error) {
	DB := getConn()
	rows, err := DB.Query("select * from orders where `Pid`=?", pid)
	if err != nil {
		log.Printf("SelectProductByPid fail: err=%v", err)
		return
	}

	for rows.Next() {
		err = rows.Scan(&product.Pid, &product.OwnerId, &product.PName, &product.Detail, &product.Remain, &product.Price, &product.PType)
	}
	return
}
