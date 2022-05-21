package entry

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type User struct {
	Uid      string
	Mail     string
	Username string
	Password string
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
	Amount int
	Time   timestamp.Timestamp
}
