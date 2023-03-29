package main

import (
	"context"
	xmlCoder "encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"sync"
	"time"

	"github.com/gorilla/schema"
)

var decoder *ReqDecoder

var cfg *Config

type weixinRequest struct {
	Signature string `json:"signature"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Echostr   string `json:"echostr"`
}

type xmlMsg struct {
	FromUserName string `xml:"FromUserName"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	MsgId        string `xml:"MsgId"`
	ToUserName   string `xml:"ToUserName"`
	CreateTime   int64  `xml:"CreateTime"`
}

type xml struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
}

var whiteList = map[string]string{
	"oNmT-0q5h-NTPQByNiGj1vVztgDU": "主人",
	"oNmT-0g7y0DSAU9__wdMHzJe4g40": "李小姐",
}

// 用于微信公众号服务器绑定通过验证
// nginx 80接口反向代理
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// 解析url的query参数
		ctx := context.Background()

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

			msg := &xmlMsg{}
			err = xmlCoder.Unmarshal(body, msg)
			if err != nil {
				log.Printf("xml.Unmarshal err=%+v", err)
				writer.Write([]byte(""))
				return
			}
			if msg.MsgId == "" {
				writer.Write([]byte(""))
				return
			}
			if msg.FromUserName == "" {
				writer.Write([]byte(""))
				return
			}

			masterName, ok := whiteList[msg.FromUserName]
			if !ok {
				writer.Write([]byte(""))
				return
			}
			if msg.MsgType != "text" {
				writer.Write([]byte(""))
				return
			}

			msgKey := fmt.Sprintf("%s_%d", msg.FromUserName, msg.CreateTime)
			if _, ok := decoder.Reply[msg.FromUserName].Load(msgKey); ok {
				writer.Write([]byte(""))
				return
			}

			content, err := AskForOpenAI(ctx, masterName, msg.Content)
			if err != nil {
				log.Printf("AskForOpenAI err=%+v", err)
				writer.Write([]byte(""))
				return
			}

			//content := "你好"
			decoder.Reply[msg.FromUserName].Store(msgKey, content)

			reply := &xml{
				ToUserName:   msg.FromUserName,
				FromUserName: msg.ToUserName,
				CreateTime:   time.Now().Unix(),
				MsgType:      "text",
				Content:      content,
			}
			replyByt, err := xmlCoder.Marshal(reply)
			if err != nil {
				log.Printf("Marshal err=%+v", err)
				writer.Write([]byte(""))
				return
			}
			log.Printf("replyByt=%s", string(replyByt))
			writer.Write(replyByt)
			return
		}

		writer.Write([]byte(wx.Echostr))
	})
	cfg = NewConfig()
	decoder = NewReqDecoder()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

type ReqDecoder struct {
	Decoder     *schema.Decoder
	FilterQuery sync.Map
	Reply       map[string]*sync.Map
}

func NewReqDecoder() *ReqDecoder {
	t := reflect.TypeOf(weixinRequest{})
	filter := sync.Map{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		filter.Store(field.Tag.Get("json"), struct{}{})
	}

	reply := make(map[string]*sync.Map)
	for key := range whiteList {
		reply[key] = &sync.Map{}
	}

	return &ReqDecoder{
		Decoder:     schema.NewDecoder(),
		FilterQuery: filter,
		Reply:       reply,
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
