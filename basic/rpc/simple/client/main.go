package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"

	"github.com/ct-zh/goLearn/basic/rpc/simple"
)

// rpc远程调用div方法
func main() {
	conn, err := net.Dial("tcp", ":13998")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	var result float64
	err = client.Call("DemoService.Div", simple.Args{
		A: 10,
		B: 3,
	}, &result)
	fmt.Println(result, err)

	err = client.Call("DemoService.Div", simple.Args{
		A: 10,
		B: 0,
	}, &result)
	fmt.Println(result, err)
}
