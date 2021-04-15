package main

import (
	"context"
	"log"

	"github.com/micro/go-micro/v2"
	proto "go-micro/demo/proto/hello"
)

type HelloServer struct{}

func (s HelloServer) SayHello(ctx context.Context, req *proto.SayRequest, res *proto.SayResponse) error {
	log.Println("client: ", req)
	res.Answer = "hello"
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("demo.hello"),
	)
	srv.Init()

	err := proto.RegisterHelloHandler(srv.Server(), new(HelloServer))
	if err != nil {
		panic(err)
	}

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
