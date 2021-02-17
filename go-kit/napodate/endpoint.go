package napodate

// 该文件包含我们的端点，这些端点会将来自客户端的请求映射到我们的内部服务

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

// 公开端点
type Endpoints struct {
	GetEndPoint      endpoint.Endpoint
	StatusEndPoint   endpoint.Endpoint
	ValidateEndPoint endpoint.Endpoint
}

// MakeGetEndpoint 返回 「get」服务的响应
func MakeGetEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(getRequest) // 我们只需要请求，不需要使用它的值
		d, err := srv.Get(ctx)
		if err != nil {
			return getResponse{
				Date: d,
				Err:  err.Error(),
			}, nil
		}
		return getResponse{
			Date: d,
			Err:  "",
		}, nil
	}
}

// MakeStatusEndpoint 返回 「status」服务的响应
func MakeStatusEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(statusRequest)
		status, err := srv.Status(ctx)
		if err != nil {
			return statusResponse{Status: status}, err
		}

		return statusResponse{Status: status}, nil
	}
}

// MakeValidateEndpoint 返回「validate」服务的响应
func MakeValidateEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(validateRequest)
		result, err := srv.Validate(ctx, req.Date)
		if err != nil {
			return validateResponse{
				Valid: result,
				Err:   err.Error(),
			}, nil
		}

		return validateResponse{
			Valid: result,
			Err:   "",
		}, nil
	}
}

// get 端点映射
//func (e Endpoints) Get(ctx context.Context) (string, error) {
//	return "", nil
//	//req := getRequest{}
//	//resp, err := e.GetEndPoint(ctx, req)
//	//if err != nil {
//	//	return "", nil
//	//}
//	//getResp := resp.(getResponse)
//	//if getResp.Err != "" {
//	//	return "", errors.New(getResp.Err)
//	//}
//	//return getResp.Date, nil
//}

// Status 端点映射
//func (e Endpoints) Status(ctx context.Context) (string, error) {
//	req := statusRequest{}
//	resp, err := e.StatusEndPoint(ctx, req)
//	if err != nil {
//		return "", err
//	}
//	statusResp := resp.(statusResponse)
//	return statusResp.Status, nil
//}

// Validate 端点映射
//func (e Endpoints) Validate(ctx context.Context, date string) (bool, error) {
//	req := validateRequest{Date: date}
//	resp, err := e.ValidateEndPoint(ctx, req)
//	if err != nil {
//		return false, err
//	}
//	validateResp := resp.(validateResponse)
//	if validateResp.Err != "" {
//		return false, errors.New(validateResp.Err)
//	}
//	return validateResp.Valid, nil
//}
