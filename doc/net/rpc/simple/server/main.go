package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/ct-zh/goLearn/doc/net/rpc/simple"
)

// 这样写的问题：不利于后期的维护和工作的切割
func main() {
	// step1. 注册方法
	_ = rpc.Register(simple.DemoService{})
	// 也可以
	//rpc.RegisterName("DemoService", new (simple.DemoService))

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
