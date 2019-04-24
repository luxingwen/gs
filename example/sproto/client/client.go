package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/luxingwen/gs/example/sproto/sp"
	"github.com/luxingwen/gs/sproto"
)

func main() {
	conn, err := net.Dial("tcp4", "127.0.0.1:5573")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("start client ", conn.RemoteAddr().Network())
	defer conn.Close()

	protocl := sproto.NewSprotoPtotocol(sp.Protocols)

	data, err := protocl.Marshal(&sp.LoginRequest{Uid: "fjkaflkajlk"})
	if err != nil {
		log.Fatal(err)
	}
	err = protocl.Write(conn, data...)
	if err != nil {
		log.Fatal(err)
	}
	b, err := protocl.Read(conn)
	if err != nil {
		log.Fatal(err)
	}
	res, err := protocl.Unmarshal(b)
	if err != nil {
		log.Fatal(err)
	}
	rdata, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("rdata:", string(rdata))
	fmt.Println("client close.")
	select {}
}
