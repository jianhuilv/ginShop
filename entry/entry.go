package entry

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type User struct {
	Uid      string `json:"uid"`
	Mail     string `json:"mail"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type PType struct {
	Sid      string `json:"sid"`
	TypeName string `json:"type_name"`
}

type Admin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Product struct {
	Pid     string `json:"pid"`
	OwnerId string `json:"owner_id"`
	PName   string `json:"p_name"`
	Detail  string `json:"detail"`
	Image   string `json:"image"`
	PType   string `json:"p_type"`
	Remain  int    `json:"remain"`
	Price   string `json:"price"`
}

type Sales struct {
	Sid    string `json:"sid"`
	Pid    string `json:"pid"`
	Remain int    `json:"remain"`
	Sold   int    `json:"sold"`
}

type Saler struct {
	Sid      string `json:"sid"`
	Password string `json:"password"`
	SName    string `json:"s_name"`
}

type Cart struct {
	Pid     string `json:"pid"`
	OwnerId string `json:"owner_id"`
	Amount  int    `json:"amount"`
}

type Order struct {
	Oid    string              `json:"oid"`
	Pid    string              `json:"pid"`
	Uid    string              `json:"uid"`
	Amount int                 `json:"amount"`
	Time   timestamp.Timestamp `json:"time"`
}
