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
	Sid      string
	TypeName string
}

type Admin struct {
	Username string
	Password string
}

type Product struct {
	Pid     string
	OwnerId string
	PName   string
	Detail  string
	Image   string
	PType   string
	Remain  int
	Price   string
}

type Sales struct {
	Sid    string
	Pid    string
	Remain int
	Sold   int
}

type Saler struct {
	Sid      string
	Password string
	SName    string
}

type Cart struct {
	Pid     string
	OwnerId string
	Amount  int
}

type Order struct {
	Oid    string
	Pid    string
	Uid    string
	Amount int
	Time   timestamp.Timestamp
}
