package main

import (
	"fmt"
	rpcdemo "github.com/simonzs/crawler_go/rpc"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	coon, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(coon)
	var result float64

	err = client.Call("DemoService.Div",
		rpcdemo.Args{A: 4, B: 3}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	err = client.Call("DemoService.Div",
		rpcdemo.Args{A: 4, B: 0}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
