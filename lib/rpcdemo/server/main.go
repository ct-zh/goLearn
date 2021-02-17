package main

import (
	"Crawler/config"
	"Crawler/rpc"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	_ = rpc.Register(rpcdemo.DemoService{})
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.RpcPort))
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
