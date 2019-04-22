package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/luxingwen/gs"
	pb "github.com/luxingwen/gs/example/proto"
)

func main() {
	ln, err := net.Listen("tcp", ":5573")
	if err != nil {
		log.Fatal(err)
	}
	list := []*gs.Protoc{
		&gs.Protoc{
			Type: 1001,
			Name: "login",
			Msg:  &pb.CsLogin{},
		},
	}
	protocl := gs.NewProtobuf(list)
	f := func() gs.ClientCallBack {
		return &User{}
	}
	server := gs.NewServer(ln, f, protocl)

	go server.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	fmt.Printf("server closing down (signal: %v)\n", sig)
	server.Stop()
}

type User struct {
	ClientId int64
}

func (u *User) OnConnect(client *gs.Client) bool {
	fmt.Println("onconnect client id = ", client.ClientId())
	return true
}

func (u *User) OnMessage(client *gs.Client, msg interface{}) {
	fmt.Println("onmessage client id = ", client.ClientId())
	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("marshal err", err)
		return
	}
	fmt.Println("data: ", string(data))
}

func (u *User) OnClose(client *gs.Client) {
	fmt.Println("on close client id = ", client.ClientId())
}
