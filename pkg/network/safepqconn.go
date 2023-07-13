package network

import (
	"context"
	"io"
	"log"
	"net"
	"time"

	"github.com/open-quantum-safe/liboqs-go/oqs"
)

var KEMProtocol = []byte("/kem/handshake/1.0.0")

func GetKEMPubKey() oqs.KeyEncapsulation {
	kemName := "Kyber512"
	kem := oqs.KeyEncapsulation{}
	if err := kem.Init(kemName, nil); err != nil {
		log.Fatal(err)
	}
	return kem
}

type SafeConnection struct {
	listener     net.Listener
	dialer       net.Dialer
	kemServer    oqs.KeyEncapsulation
	clientPubKey []byte
	disconnect   chan struct{}
	Health       chan bool
}

func (sc *SafeConnection) Init(r io.Reader, w io.Writer) *SafeConnection {

	kemServer := GetKEMPubKey()
	sc.kemServer = kemServer
	sc.listener, _ = net.Listen("tcp", ":22456")
	sc.disconnect = make(chan struct{})
	sc.Health = make(chan bool)
	return sc
}

var Connected []*SafeConnection

func SafeConnect(addr string) error {
	var d net.Dialer
	var conn net.Conn
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	var err error
	conn, err = d.DialContext(ctx, "tcp", addr)
	if err != nil {
		return err
	}

	defer conn.Close()

	return nil

}

func (sc *SafeConnection) AddPeerAndDial(addr string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	sc.dialer.DialContext(ctx, "tcp", addr)
}
