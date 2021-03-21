package simple

import "errors"

// Service.Method 远程调用的服务方法

type DemoService struct{}

type Args struct {
	A, B int
}

// result must be a pointer type
func (DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}

	*result = float64(args.A) / float64(args.B)
	return nil
}
