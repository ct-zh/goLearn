package simple

import "errors"

// DemoService Service.Method 远程调用的服务方法
type DemoService struct{}

type Args struct {
	A, B int
}

// Div result must be a pointer type
// 必须满足Go语言的RPC规则：方法只能有两个可序列化的参数，其中第二个参数是指针类型，并且返回一个error类型，同时必须是公开的方法。
func (DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}

	*result = float64(args.A) / float64(args.B)
	return nil
}

//
