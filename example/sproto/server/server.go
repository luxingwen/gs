package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/luxingwen/gs"
	"github.com/luxingwen/gs/example/sproto/sp"
	"github.com/luxingwen/gs/sproto"
)

func main() {
	ln, err := net.Listen("tcp", ":5573")
	if err != nil {
		log.Fatal(err)
	}

	protocl := sproto.NewSprotoPtotocol(sp.Protocols)
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
	spMsg := msg.(sproto.Sproto)

	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("marshal err", err)
		return
	}
	fmt.Println("data: ", string(data))

	switch spMsg.GetType() {
	case 10001:
		res, err := u.LoginRequest(spMsg.(*sp.LoginRequest))
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		if res != nil {
			client.WriteMsg(res)
		}
	default:
	}
}

func (u *User) OnClose(client *gs.Client) {
	fmt.Println("on close client id = ", client.ClientId())
}

func (u *User) LoginRequest(req *sp.LoginRequest) (r interface{}, err error) {
	fmt.Println("login uid:", req.Uid)
	user := &sp.UserInfo{
		Uid:    req.Uid,
		Number: 123,
		Name:   "shuaishaui",
		Head:   "1234",
	}
	r = &sp.LoginResponse{Info: user}
	return
}
