# 爬虫
1. `resp, err := http.Get(url)`,返回的response记得close: `defer resp.Body.Close()`
2. `all, err := ioutil.ReadAll(resp.Body)`，读取当前页面内容
3. `resp.StatusCode`用来判断请求状态

## 编码转换：
下载两个包:  编码转换`gopm get -g -v golang.org/x/text`；自动检测编码：`gopm get -g -v golang.org/x/net/html`
```go

// 自动检测编码
func determineEncoding(r io.Reader) encoding.Encoding {
    bytes, err := bufio.NewReader(r).Peek(1024)
    if err != {
        panic(err)
    }
    e, _, _ := charset.DetermineEncoding(bytes, "")
    return e
}

// gbk转utf8
utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
```

## 正则表达式
这里pattern通常使用``符号来包围
```go
re := regexp.MustCompile(pattern)
re.FindString(text)
re.FindAllString(text)
```

可以用括号单独拉出来需要匹配的子项,比如在某个文本里面匹配以下内容，并且我还需要获取href和a标签里面的内容
```go
// <a href="http://www.zhenai.com/zhenghun/guangzhou" data-v-2cb5b6a2>广州</a>
re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)" [^>]*>([^<]+)</a>`)
matches := re.FindAllSubmatch(contents, -1)
for _, m := range matches {
    fmt.Printf("City: %s, Url: %s\n", m[2], m[1])
}
fmt.Printf("Matches found: %d\n", len(matches))
```




## 单任务爬虫架构设计
按照 http://www.zhenai.com/zhenghun 这个列表解析
分为三层： 城市列表解析器，城市解析器，用户解析器；
解析器Parser:
1. 先将内容转成utf8
2. 输出request{url , Parser}，和 Item列表
架构：
1. seed采集到url列表 -> engine -> 加到任务队列
2. engine 取出队列，传递url -> Fetcher 做http请求，将数据回传给engine
3. engine data -> parser
4. parser 解析出requests, items -> engine


## 生成error
1. `errors.New(text)`
2. `fmt.Errorf()`


# 分布式系统：
1. 多个节点；
2. 消息传递；
3. 完成特定需求

分布式架构与微服务架构：分布式是指导节点之间如何通信；微服务是鼓励按照业务来划分模块；微服务架构可以通过分布式架构来实现。
微服务通常要配合自动化测试，部署，服务发现等。

## 多个节点
容错性；可扩展性（性能）；固有分布性；

## 一般消息传递的方法：
对外： REST
模块内部： RPC
模块之间：中间件，REST

RPC: jsonrpc,grpc,Thrift

### rpc的序列化
因为rpc相当于直接对方法进行远程调用，如果参数是整数/字符串之类的变量倒是可以直接传输，但如果调用的是一个闭包或者其他复杂的内容就不能进行传输；所以一般需要提供一个将参数转换为json字符串的方法，以及对应的解析方法。这个我们称之为rpc的序列化。

## go语言实现rpcDemo
rpc服务器：
```go
package main

import (
	"Crawler/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(rpcdemo.DemoService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", conn)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}
```

调用的方法：
```go
package rpcdemo

import "errors"

// Service.Method
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
```

测试方法：先开启main服务，然后`telnet localhost 1234`,输入`{"method":"DemoService.Div","params":[{"A":3,"B":4}],"id":1}`即可获取到结果；

