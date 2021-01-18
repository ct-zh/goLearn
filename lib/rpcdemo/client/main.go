package main

import (
	"Crawler/config"
	rpcdemo "Crawler/rpc"
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", config.RpcPort))
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	var result float64
	err = client.Call("DemoService.Div", rpcdemo.Args{
		A: 10,
		B: 3,
	}, &result)
	fmt.Println(result, err)

	err = client.Call("DemoService.Div", rpcdemo.Args{
		A: 10,
		B: 0,
	}, &result)
	fmt.Println(result, err)

}
