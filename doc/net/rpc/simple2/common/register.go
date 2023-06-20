package common

import (
	"fmt"
	"net/rpc"
)

const DemoServiceName = "github.com/ct-zh/goLearn/doc/net/rpc/simple2/common/DemoService"

// DemoServiceInterface 定义接口规范
type DemoServiceInterface interface {
	Div(args Args, result *float64) error
}

func RegisterDemoService(svc DemoServiceInterface) error {
	return rpc.RegisterName("DemoService", svc)
}

type DemoServiceClient struct {
	*rpc.Client
}

// 规范实现interface的方法
var _ DemoServiceInterface = (*DemoServiceClient)(nil)

func (d *DemoServiceClient) Div(args Args, result *float64) error {
	fmt.Printf("args = %+v", args)
	return d.Client.Call(DemoServiceName+".Div", args, result)
	//return d.Client.Call("DemoService.Div", args, result)
}

// NewDemoClient 初始化client service
func NewDemoClient(network, address string) (*DemoServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &DemoServiceClient{Client: c}, nil
}
