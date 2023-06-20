package main

import (
	"log"
	"net"
	"net/rpc/jsonrpc"

	"github.com/ct-zh/goLearn/doc/net/rpc/simple2/common"
)

func main() {
	// step1. 这样不仅可以避免命名服务名称的工作，同时也保证了传入的服务对象满足了RPC接口的定义
	common.RegisterDemoService(new(common.DemoService))

	// step2. 启动服务监听
	listener, err := net.Listen("tcp", ":13998")
	if err != nil {
		panic(err)
	}

	for {
		// step3. 处理请求，Accept获取请求链接
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", conn)
			continue
		}
		// step4.处理请求
		go jsonrpc.ServeConn(conn)
	}
}
