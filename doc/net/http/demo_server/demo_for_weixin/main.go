package main

import (
	"fmt"
	"github.com/gorilla/schema"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"sync"
)

var decoder *ReqDecoder

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
		// 解析url的query参数
		log.Printf("query=%+v", request.URL.Query())
		wx, err := decoder.Decode(request.URL.Query())
		if err != nil {
			log.Printf("decoder.Decode err=%+v", err)
			writer.Write([]byte(""))
			return
		}
		log.Printf("wx=%+v", wx)

		if request.Method == http.MethodPost {
			body, err := ioutil.ReadAll(request.Body)
			if err != nil {
				fmt.Printf("err= %+v \n", err)
				return
			}
			log.Printf("request= %s \n", string(body))
		}

		writer.Write([]byte(wx.Echostr))
	})
	decoder = NewReqDecoder()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

type ReqDecoder struct {
	Decoder     *schema.Decoder
	FilterQuery sync.Map
}

func NewReqDecoder() *ReqDecoder {
	t := reflect.TypeOf(weixinRequest{})
	filter := sync.Map{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		filter.Store(field.Tag.Get("json"), struct{}{})
	}
	return &ReqDecoder{
		Decoder:     schema.NewDecoder(),
		FilterQuery: filter,
	}
}

func (r *ReqDecoder) Decode(val url.Values) (*weixinRequest, error) {
	filterQuery := make(map[string][]string)
	for key, strings := range val {
		if _, ok := r.FilterQuery.Load(key); ok {
			filterItem := strings
			filterQuery[key] = filterItem
		}
	}
	log.Printf("filter query=%+v", filterQuery)
	wx := &weixinRequest{}
	err := r.Decoder.Decode(wx, filterQuery)
	if err != nil {
		return nil, err
	}
	return wx, nil
}
