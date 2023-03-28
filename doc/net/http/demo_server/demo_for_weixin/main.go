package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
)

type weixinRequest struct {
	Signature string `json:"signature"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Echostr   string `json:"echostr"`
}

// 用于微信公众号服务器绑定通过验证
// nginx 80接口反向代理
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		wx := &weixinRequest{}
		decoder := schema.NewDecoder()
		err := decoder.Decode(wx, request.URL.Query())
		if err != nil {
			fmt.Printf("err= %+v \n", err)
			return
		}
		writer.Write([]byte(wx.Echostr))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
