package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"

	"github.com/ct-zh/goLearn/doc/net/rpc/simple"
)

// rpc远程调用div方法
func main() {
	// step1. 申明服务器链接
	conn, err := net.Dial("tcp", ":13998")
	if err != nil {
		panic(err)
	}

	// step2. 初始化client
	client := jsonrpc.NewClient(conn)

	// step3. 使用call发起请求
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
