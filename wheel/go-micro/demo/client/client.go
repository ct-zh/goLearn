package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	demo_hello "go-micro/demo/proto/hello"
)

func main() {
	srv := micro.NewService(micro.Name("demo.hello.client"))
	srv.Init()

	hello := demo_hello.NewHelloService("demo.hello", srv.Client())

	res, err := hello.SayHello(context.TODO(),
		&demo_hello.SayRequest{Message: "你好你好"})

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
