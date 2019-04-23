package gs

import (
	"encoding/binary"
	"errors"
	"io"
	"net"
	"reflect"

	"github.com/golang/protobuf/proto"
)

var defaultSize uint32 = 1024 * 16

type Protoc struct {
	Type uint16
	Name string
	Msg  interface{}
}

type Protobuf struct {
	idMap map[uint16]*Protoc
	msgId map[reflect.Type]uint16
}

func NewProtobuf(list []*Protoc) *Protobuf {
	p := &Protobuf{idMap: make(map[uint16]*Protoc, 0), msgId: make(map[reflect.Type]uint16)}
	for _, item := range list {
		if _, ok := p.idMap[item.Type]; ok {
			panic("allready register protoc")
		}
		p.idMap[item.Type] = item
		p.msgId[reflect.TypeOf(item.Msg)] = item.Type
	}
	return p
}

func (p *Protobuf) Read(conn net.Conn) (r []byte, err error) {
	bufMsgLen := make([]byte, 4)

	// read len
	if _, err := io.ReadFull(conn, bufMsgLen); err != nil {
		return nil, err
	}

	// parse len
	var msgLen uint32
	msgLen = binary.BigEndian.Uint32(bufMsgLen)

	// check len
	if msgLen > defaultSize {
		return nil, errors.New("message too long")
	} else if msgLen < 1 {
		return nil, errors.New("message too short")
	}

	// data
	msgData := make([]byte, msgLen)
	if _, err := io.ReadFull(conn, msgData); err != nil {
		return nil, err
	}

	return msgData, nil
}

func (p *Protobuf) Write(conn net.Conn, args ...[]byte) error {
	// get len
	var msgLen uint32
	for i := 0; i < len(args); i++ {
		msgLen += uint32(len(args[i]))
	}

	// check len
	if msgLen > defaultSize {
		return errors.New("message too long")
	} else if msgLen < 1 {
		return errors.New("message too short")
	}

	msg := make([]byte, 4+msgLen)

	binary.BigEndian.PutUint32(msg, msgLen)

	// write data
	l := 4
	for i := 0; i < len(args); i++ {
		copy(msg[l:], args[i])
		l += len(args[i])
	}
	conn.Write(msg)
	return nil
}

func (p *Protobuf) Unmarshal(data []byte) (interface{}, error) {
	if len(data) < 4 {
		return nil, errors.New("protobuf data too short")
	}

	// id
	var id uint16
	id = binary.BigEndian.Uint16(data)

	// msg
	i := p.idMap[id]

	msg := reflect.New(reflect.TypeOf(i.Msg).Elem()).Interface()
	if err := proto.UnmarshalMerge(data[2:], msg.(proto.Message)); err != nil {
		return nil, err
	}
	return &Protoc{Type: i.Type, Name: i.Name, Msg: msg}, nil

}

func (p *Protobuf) Marshal(msg interface{}) ([][]byte, error) {
	msgId, ok := p.msgId[reflect.TypeOf(msg)]
	if !ok {
		return nil, errors.New("not register protobuf")
	}
	id := make([]byte, 2)
	binary.BigEndian.PutUint16(id, msgId)

	// data

	data, err := proto.Marshal(msg.(proto.Message))
	return [][]byte{id, data}, err
}
