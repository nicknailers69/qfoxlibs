package core

import (
	"net"
)

type QFoxCore interface {
	NetworkID() string
	Version() string
	Connection() (*net.Conn, error)
}

type Server struct {
	node Node
}

type QCore struct {
	S *Server
}
