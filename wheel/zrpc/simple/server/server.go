package main

import (
	"context"
	"flag"
	"log"

	"github.com/ct-zh/goLearn/wheel/zrpc/simple/pb"

	grpc "google.golang.org/grpc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
}

type Hello struct {
	pb.UnimplementedGreeterServer
}

func (h *Hello) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello " + request.Name}, nil
}

var cfgFile = flag.String("f", "./server/hello.yaml", "cfg file")

func main() {
	flag.Parse()

	var cfg Config
	conf.MustLoad(*cfgFile, &cfg)

	srv, err := zrpc.NewServer(cfg.RpcServerConf, func(server *grpc.Server) {
		pb.RegisterGreeterServer(server, &Hello{})
	})
	if err != nil {
		log.Fatal(err)
	}
	srv.Start()
}
