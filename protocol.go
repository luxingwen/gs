package gs

import (
	"net"
)

type Protocol interface {
	Read(net.Conn) ([]byte, error)
	Marshal(msg interface{}) ([][]byte, error)
	Unmarshal(data []byte) (interface{}, error)
	Write(net.Conn, ...[]byte) error
}
