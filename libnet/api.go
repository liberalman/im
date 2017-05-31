package libnet

import (
	"io"
	"net"
	"time"
)

type Protocol interface {
	NewCodec(rw io.ReadWriter) Codec
}

type Codec interface {
	Receive() ([]byte, error)
	Send(interface{}) error
	Close() error
}

func Serve(network, address string, protocol Protocol, sendChanSize int) (*Server, error) {
	listener, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}
	return NewServer(listener, protocol, sendChanSize), nil
}

func Connect(network, address string, protocol Protocol, sendChanSize int) (*Session, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return NewSession(protocol.NewCodec(conn), sendChanSize), nil
}

func ConnectTimeout(network, address string, timeout time.Duration, protocol Protocol, sendChanSize int) (*Session, error) {
	conn, err := net.DialTimeout(network, address, timeout)
	if err != nil {
		return nil, err
	}
	return NewSession(protocol.NewCodec(conn), sendChanSize), nil
}
