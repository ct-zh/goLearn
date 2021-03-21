package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/LannisterAlwaysPaysHisDebts/goLearn/basic/rpc/simple"
)

func main() {
	// 注册方法
	_ = rpc.Register(simple.DemoService{})

	// 启动服务监听
	listener, err := net.Listen("tcp", ":13998")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", conn)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}
