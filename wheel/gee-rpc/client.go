package geerpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"geerpc/codec"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

// 需要客户端处理超时的地方有：
//
// 1. 与服务端建立连接，导致的超时;
// 2. 发送请求到服务端，写报文导致的超时
// 3. 等待服务端处理时，等待处理导致的超时（比如服务端已挂死，迟迟不响应）
// 4. 从服务端接收响应时，读报文导致的超时

// 客户端
type Client struct {
	cc      codec.Codec      // 编码器
	opt     *Option          // 协商opt
	sending sync.Mutex       // 为了保证请求的有序发送，即防止出现多个请求报文混淆。
	header  codec.Header     // 只有在请求发送时才需要，而请求发送是互斥的，因此每个客户端只需要一个，声明在 Client 结构体中可以复用。
	mu      sync.Mutex       // 结构体锁，对结构体的操作都需要加锁
	seq     uint64           // 请求编号
	pending map[uint64]*Call // 存储未处理完的请求，键是编号，值是 Call 实例

	closing  bool // 客户端关闭
	shutdown bool // 服务端关闭
}

type clientResult struct {
	client *Client
	err    error
}

// 这种写法保证 client已经完成了io.Closer的全部接口
var _ io.Closer = (*Client)(nil)

// 定义错误信息
var ErrShutdown = errors.New("connection is shut down")

// 创建连接
// @network 网络格式
// @address 地址
// @opts option参数
func Dial(network, address string, opts ...*Option) (client *Client, err error) {
	return dialTimeout(NewClient, network, address, opts...)
}

type newClientFunc func(conn net.Conn, opt *Option) (client *Client, err error)

func dialTimeout(f newClientFunc, network, address string, opts ...*Option) (client *Client, err error) {
	opt, err := parseOptions(opts...)
	if err != nil {
		return nil, err
	}

	// 超时处理： 如果连接创建超时，将返回错误
	conn, err := net.DialTimeout(network, address, opt.ConnectTimeout)
	if err != nil {
		return nil, err
	}

	// close the connection if client is nil
	defer func() {
		if err != nil {
			_ = conn.Close()
		}
	}()

	// 使用子协程执行 NewClient , 主协程增加计时器
	ch := make(chan clientResult)
	go func() {
		client, err := f(conn, opt)
		ch <- clientResult{
			client: client,
			err:    err,
		}
	}()

	// 为0代表不计超时时间
	if opt.ConnectTimeout == 0 {
		result := <-ch
		return result.client, result.err
	}

	select {
	case <-time.After(opt.ConnectTimeout): // 超时逻辑
		return nil, fmt.Errorf("rpc client: connect timeout: expect within %s", opt.ConnectTimeout)
	case result := <-ch:
		return result.client, result.err
	}
}

// 解析opt参数, 拼凑出option
func parseOptions(opts ...*Option) (*Option, error) {
	if len(opts) == 0 || opts[0] == nil {
		return DefaultOption, nil
	}
	if len(opts) != 1 {
		return nil, errors.New("number of options is more than 1")
	}

	//
	opt := opts[0]
	opt.MagicNumber = DefaultOption.MagicNumber
	if opt.CodecType == "" {
		opt.CodecType = DefaultOption.CodecType
	}

	return opt, nil
}

// 根据conn连接初始化客户端, 并创建协程接收数据
func NewClient(conn net.Conn, opt *Option) (*Client, error) {
	// 获取编码器
	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		err := fmt.Errorf("invalid codec type %s", opt.CodecType)
		log.Println("rpc client: codec error:", err)
		return nil, err
	}

	// 首先想服务端发送option， 协商好编码方式
	if err := json.NewEncoder(conn).Encode(opt); err != nil {
		log.Println("rpc client: options error: ", err)
		_ = conn.Close()
		return nil, err
	}

	// 初始化client
	client := &Client{
		seq:     1, // 请求id从1开始
		cc:      f(conn),
		opt:     opt,
		pending: make(map[uint64]*Call),
	}
	go client.receive() // 创建子协程接收响应

	return client, nil
}

// 实现 io.Closer;将client.closing 加锁设置为true
func (client *Client) Close() error {
	client.mu.Lock()
	defer client.mu.Unlock()
	if client.closing {
		return ErrShutdown
	}
	client.closing = true
	return client.cc.Close()
}

// 判断是否可用
func (client *Client) IsAvailable() bool {
	client.mu.Lock()
	defer client.mu.Unlock()
	return !client.shutdown && !client.closing
}

// 协程接收消息; 接收到的响应有三种情况:
// 1. call 不存在，可能是请求没有发送完整，或者因为其他原因被取消，但是服务端仍旧处理了。
// 2. call 存在，但服务端处理出错，即 h.Error 不为空。
// 3. call 存在，服务端处理正常，那么需要从 body 中读取 Reply 的值。
func (client *Client) receive() {
	var err error

	for err == nil {
		var h codec.Header

		// readheader是从 conn中读出数据并写到h里面
		if err = client.cc.ReadHeader(&h); err != nil {
			break
		}

		// 获得的是远程调用返回的结果， 先根据header的seq获取对应的call请求
		call := client.removeCall(h.Seq)

		switch {
		case call == nil: // call不存在
			err = client.cc.ReadBody(nil)

		case h.Error != "": // header.Error 出现错误
			call.Error = fmt.Errorf(h.Error)
			err = client.cc.ReadBody(nil)
			call.done() // 申明call已经处理完毕了,结束单次call调用请求;

		default: // 正常获取调用返回结果

			// 从 conn中读出数据写到 call.reply 里面
			err = client.cc.ReadBody(call.Reply)
			if err != nil {
				call.Error = errors.New("reading body " + err.Error())
			}
			call.done()
		}
	}

	// 出现error， 关闭所有calls
	client.terminateCalls(err)
}
