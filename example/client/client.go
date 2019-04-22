package main

import (
	"fmt"
	"log"
	"net"

	"github.com/luxingwen/gs"
	pb "github.com/luxingwen/gs/example/proto"
)

func main() {
	conn, err := net.Dial("tcp4", "127.0.0.1:5573")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("start client ", conn.RemoteAddr().Network())
	defer conn.Close()
	list := []*gs.Protoc{
		&gs.Protoc{
			Type: 1001,
			Name: "login",
			Msg:  &pb.CsLogin{},
		},
	}
	protocl := gs.NewProtobuf(list)

	var id string = "dhjfkaflk"
	data, err := protocl.Marshal(&pb.CsLogin{Id: &id})
	if err != nil {
		log.Fatal(err)
	}
	err = protocl.Write(conn, data...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("client close.")
	select {}
}
