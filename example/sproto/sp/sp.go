package sp

import (
	"reflect"

	"github.com/luxingwen/gs/sproto"
)

type LoginRequest struct {
	Uid string `sproto:"string,0,name=uid"`
}
type LoginResponse struct {
	Errcode int64     `sproto:"integer,0,name=errcode"`
	Info    *UserInfo `sproto:"struct,1,name=info"`
}

func (*LoginRequest) GetType() uint16 {
	return uint16(10001)
}
func (*LoginRequest) GetMode() uint8 {
	return uint8(0)
}
func (*LoginResponse) GetType() uint16 {
	return uint16(10001)
}
func (*LoginResponse) GetMode() uint8 {
	return uint8(1)
}

type RegisterRequest struct {
	Uid  string `sproto:"string,0,name=uid"`
	Name string `sproto:"string,1,name=name"`
	Head string `sproto:"string,2,name=head"`
}
type RegisterResponse struct {
	Errcode int64     `sproto:"integer,0,name=errcode"`
	Info    *UserInfo `sproto:"struct,1,name=info"`
}

func (*RegisterRequest) GetType() uint16 {
	return uint16(10002)
}
func (*RegisterRequest) GetMode() uint8 {
	return uint8(0)
}
func (*RegisterResponse) GetType() uint16 {
	return uint16(10002)
}
func (*RegisterResponse) GetMode() uint8 {
	return uint8(1)
}

type UserInfo struct {
	Uid     string `sproto:"string,0,name=uid"`
	Number  int64  `sproto:"integer,1,name=number"`
	Name    string `sproto:"string,2,name=name"`
	Head    string `sproto:"string,3,name=head"`
	Gold    int64  `sproto:"integer,4,name=gold"`
	Diamond int64  `sproto:"integer,5,name=diamond"`
	Level   int64  `sproto:"integer,6,name=level"`
	Exp     int64  `sproto:"integer,7,name=exp"`
	Energy  int64  `sproto:"integer,8,name=energy"`
}

var Protocols []*sproto.SpProtocol = []*sproto.SpProtocol{
	&sproto.SpProtocol{
		Type:     10001,
		Name:     "login",
		Request:  reflect.TypeOf(&LoginRequest{}),
		Response: reflect.TypeOf(&LoginResponse{}),
	},
	&sproto.SpProtocol{
		Type:     10002,
		Name:     "register",
		Request:  reflect.TypeOf(&RegisterRequest{}),
		Response: reflect.TypeOf(&RegisterResponse{}),
	},
}
