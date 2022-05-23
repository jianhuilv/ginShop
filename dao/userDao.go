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
		log.Printf("InsertIntoCart fail: res=%v,err=%w", res, err)
		return
	}
	return
}

func InsertOrder(order entry.Order) (success bool, err error) {
	DB := getConn()
	rows, err := DB.Prepare("insert into order(`oid`,`pid`,`amount`,`status`) values (?,?,?,?,?)")
	if err != nil {
		log.Printf("InsertIntoCart fail: err=%w", err)
		return
	}
	res, err := rows.Exec(order.Oid, order.Pid, order.Amount, "尚未付款")
	if err != nil {
		log.Printf("InsertIntoCart fail: res=%v,err=%w", res, err)
		return
	}
	return
}

func UpDateStatusOfOrder(order entry.Order) (success bool, err error) {
	return
}

func SelectOrdersByUid(uid string) (orders []entry.Order, err error) {
	return
}

func SelectProductByPid(pid string) (product entry.Product, err error) {
	return
}
