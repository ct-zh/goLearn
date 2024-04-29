package main

import (
	"context"
	"fmt"
	"net/http"
)

type requestIdType int

var requestIDKey requestIdType = 0

// 下面例子创建了一个ValueCtx, 并将当次请求header中的X-Request-ID存储在其中
// 在业务逻辑里可以直接调用ctx.Value获取X-Request-ID的值
func main() {
	// 创建一个简单的 HTTP 处理器
	handler := WithRequestID(http.HandlerFunc(Handle))

	// 开启监听服务
	err := http.ListenAndServe(":8085", handler)
	if err != nil {
		panic(err)
	}
}

// 创建一个 context，并存储header中的X-Request-ID
func WithRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			// 从 header 中提取 request-id
			reqID := req.Header.Get("X-Request-ID")

			// 创建 valueCtx。其中requestIDKey使用自定义的类型，不容易冲突
			ctx := context.WithValue(req.Context(), requestIDKey, reqID)

			req = req.WithContext(ctx)

			// 调用 HTTP 处理函数
			next.ServeHTTP(rw, req)
		})
}

// 获取 request-id
func GetRequestID(ctx context.Context) string {
	return ctx.Value(requestIDKey).(string)
}

func Handle(_ http.ResponseWriter, req *http.Request) {
	// 拿到 reqId，后面可以记录日志等等
	reqID := GetRequestID(req.Context())
	fmt.Println("get reqId = ", reqID)
}
