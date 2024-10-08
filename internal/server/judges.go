package gameserver

import (
	"net"
)

type Judge struct {
	Name             string
	ControllerClient net.Conn
}

func NewJudge(name string, tcpConn net.Conn) *Judge {
	return &Judge{
		Name:             name,
		ControllerClient: tcpConn,
	}
}
