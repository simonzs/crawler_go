package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	rpcdemo "github.com/simonzs/crawler_go/rpc"
)

/*
telnet locahost 123
{"method": "DemoService.Div", "params":[{"A":3, "B":4}], "id":1}
*/

func main() {
	rpc.Register(rpcdemo.DemoService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
