package main

import (
	"fmt"
	"github.com/gorilla/schema"
	"net/http"
)

type weixinRequest struct {
	Signature string `json:"signature"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Echostr   string `json:"echostr"`
}

// 用于微信公众号服务器绑定通过验证
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
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
