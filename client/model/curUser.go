package model

import "net"

type CurUser struct {
	Conn net.Conn
	User
}
