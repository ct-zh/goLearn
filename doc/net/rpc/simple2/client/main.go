package main

import (
	"fmt"
	"github.com/ct-zh/goLearn/doc/net/rpc/simple2/common"
)

// 客户端不用再担心RPC方法名字或参数类型不匹配等低级错误的发生。
func main() {
	client, err := common.NewDemoClient("tcp", ":13998")
	if err != nil {
		panic(err)
	}
	var result *float64
	err = client.Div(common.Args{
		A: 3,
		B: 6,
	}, result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
