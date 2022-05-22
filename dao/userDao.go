package dao

import (
	"goShop/entry"
)

func NewUser(user *entry.User) {
	return
}

func SelectUidByUsernameAndPwd(user *entry.User) (uuid string) {
	return
}

func InsertIntoCart(cart entry.Cart) (err error) {
	return
}

func InsertOrder() {

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
