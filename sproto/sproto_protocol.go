package sproto

import (
	"encoding/binary"
	"errors"
	"io"
	"net"
	"reflect"
	"sync"
)

const (
	RequestModel uint8 = iota
	ResponseModel
)

var defaultSize uint32 = 1024 * 16

type SpProtocol struct {
	Type     int32
	Name     string
	Request  reflect.Type
	Response reflect.Type
}

func (p *SpProtocol) HasRequest() bool {
	return p.Request != nil
}

func (p *SpProtocol) HasResponse() bool {
	return p.Response != nil
}

type Sproto interface {
	GetType() uint16
	GetMode() uint8
}

type SprotoPtotocol struct {
	rw        *sync.RWMutex
	mProtocol map[uint16]*SpProtocol
}

func NewSprotoPtotocol(list []*SpProtocol) *SprotoPtotocol {
	s := &SprotoPtotocol{
		rw:        new(sync.RWMutex),
		mProtocol: make(map[uint16]*SpProtocol, 0),
	}
	for _, item := range list {
		if _, ok := s.mProtocol[uint16(item.Type)]; ok {
			panic("allready register sproto " + item.Name)
		}
		s.mProtocol[uint16(item.Type)] = item
	}
	return s
}

func (s *SprotoPtotocol) getSpProtocol(t uint16) (r *SpProtocol, err error) {
	s.rw.RLock()
	defer s.rw.RUnlock()
	if _, ok := s.mProtocol[t]; !ok {
		err = errors.New("not found spProtocol.")
		return
	}
	r = s.mProtocol[t]
	return
}

func (s *SprotoPtotocol) Read(conn net.Conn) (r []byte, err error) {
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

func (p *SprotoPtotocol) Write(conn net.Conn, args ...[]byte) error {
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

func (p *SprotoPtotocol) Unmarshal(data []byte) (interface{}, error) {
	if len(data) < 3 {
		return nil, errors.New("sproto data too short")
	}

	// id
	var id uint16
	id = binary.BigEndian.Uint16(data)

	sp, err := p.getSpProtocol(id)
	if err != nil {
		return nil, err
	}

	model := uint8(data[2])
	rdata, err := Unpack(data[3:])
	if err != nil {
		return nil, err
	}
	if model == RequestModel {
		reqData := reflect.New(sp.Request.Elem()).Interface()
		if _, err = Decode(rdata, reqData); err != nil {
			return nil, err
		}
		return reqData, nil
	} else {
		respData := reflect.New(sp.Response.Elem()).Interface()
		if _, err = Decode(rdata, respData); err != nil {
			return nil, err
		}
		return respData, nil
	}
	return nil, errors.New("unmarsha data err.")
}

func (p *SprotoPtotocol) Marshal(msg interface{}) ([][]byte, error) {
	sp := msg.(Sproto)
	mode := sp.GetMode()
	id := make([]byte, 3)
	binary.BigEndian.PutUint16(id, sp.GetType())
	id[2] = byte(mode)
	data, err := Encode(msg)
	if err != nil {
		return nil, err
	}
	return [][]byte{id, Pack(data)}, nil
}
