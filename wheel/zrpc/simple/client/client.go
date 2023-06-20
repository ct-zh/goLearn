package main

import (
	"context"
	"github.com/ct-zh/goLearn/wheel/zrpc/simple/pb"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	"log"
)

func main() {
	client := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "hello.rpc",
		},
	})

	conn := client.Conn()
	hello := pb.NewGreeterClient(conn)
	reply, err := hello.SayHello(context.Background(), &pb.HelloRequest{Name: "ikun"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(reply.Message)
}
