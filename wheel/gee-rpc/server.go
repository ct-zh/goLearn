package geerpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
	"strings"
	"sync"
	"time"

	"geerpc/codec"
)

// 用于确认这是一个gee rpc请求
const MagicNumber = 0x3bef5c

// 协商数据包 用于server 与client创建连接后协程数据数据格式等内容， 默认用json编码
type Option struct {
	MagicNumber int        // 用来确认这是一个geerpc请求
	CodecType   codec.Type // client的编码类型

	// 超时策略
	ConnectTimeout time.Duration // 连接超时时间， 0代表无限制
	HandleTimeout  time.Duration // 默认值为 0，即不设限
}

// 这是一个默认的Option
var DefaultOption = &Option{
	MagicNumber:    MagicNumber,
	CodecType:      codec.GobType,
	ConnectTimeout: time.Second * 10,
}

// rpc server，提供若干method
type Server struct {
	serviceMap sync.Map
}

// server构造函数
func NewServer() *Server {
	return &Server{}
}

// 创建一个默认的server, 类似于php的静态调用
var DefaultServer = NewServer()

// Accept 监听请求， 并将conn传入协程中
func (server *Server) Accept(lis net.Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println("rpc server: accept error:", err)
			return
		}
		go server.ServeConn(conn)
	}
}

// 静态调用Accept
func Accept(lis net.Listener) { DefaultServer.Accept(lis) }

// 注册服务
func (server *Server) Register(rcvr interface{}) error {
	s := newService(rcvr)
	if _, dup := server.serviceMap.LoadOrStore(s.name, s); dup {
		return errors.New("rpc: service already defined: " + s.name)
	}
	return nil
}

// 静态调用Register
func Register(rcvr interface{}) error { return DefaultServer.Register(rcvr) }

// 解析ServiceMethod为 service与method；
func (server *Server) findService(serviceMethod string) (svc *service, mtype *methodType, err error) {
	// 获取 . 的位置
	dot := strings.LastIndex(serviceMethod, ".")
	if dot < 0 {
		err = errors.New("rpc server: service/method request ill-formed: " + serviceMethod)
		return
	}

	// “Service.Method”，将其分割成 2 部分，第一部分是 Service 的名称，第二部分即方法名
	serviceName, methodName := serviceMethod[:dot], serviceMethod[dot+1:]

	// 从serviceMap里加载对应的服务
	svci, ok := server.serviceMap.Load(serviceName)
	if !ok {
		err = errors.New("rpc server: can't find service " + serviceName)
		return
	}

	svc = svci.(*service)
	mtype = svc.method[methodName]
	if mtype == nil {
		err = errors.New("rpc server: can't find method " + methodName)
	}

	return
}

// 处理单个连接请求，协商option
func (server *Server) ServeConn(conn io.ReadWriteCloser) {
	defer func() { _ = conn.Close() }()

	// 验证option,拿到编码器
	var opt Option
	if err := json.NewDecoder(conn).Decode(&opt); err != nil {
		log.Println("rpc server: options error: ", err)
		return
	}

	if opt.MagicNumber != MagicNumber {
		log.Printf("rpc server: invalid magic number %x", opt.MagicNumber)
		return
	}

	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		log.Printf("rpc server: invalid codec type %s", opt.CodecType)
		return
	}

	server.serveCodec(f(conn), &opt)
}

// 空返回体
var invalidRequest = struct{}{}

//
func (server *Server) serveCodec(cc codec.Codec, opt *Option) {
	sending := new(sync.Mutex) // 发送锁，防止同一时间有多个协程发送消息
	wg := new(sync.WaitGroup)  // 保证多协程运行完毕
	for {
		// 死循环读请求，每次请求都开启一个协程去处理
		req, err := server.readRequest(cc)
		if err != nil {
			if req == nil {
				break // 关闭请求
			}
			req.h.Error = err.Error() // 非关闭请求，可能是报错请求,
			server.sendResponse(cc, req.h, invalidRequest, sending)
			continue
		}

		// 开启异步调用逻辑
		wg.Add(1)
		go server.handleRequest(cc, req, sending, wg, opt.HandleTimeout)
	}
	wg.Wait()      // 等待所有协程运行完毕
	_ = cc.Close() // 关闭连接
}

// 请求体
type request struct {
	h            *codec.Header // header of request
	argv, replyv reflect.Value // 参数与回复消息
	mtype        *methodType
	svc          *service
}

// 读header, 处理可能是EOF的条件
func (server *Server) readRequestHeader(cc codec.Codec) (*codec.Header, error) {
	var h codec.Header
	if err := cc.ReadHeader(&h); err != nil {
		if err != io.EOF && err != io.ErrUnexpectedEOF {
			log.Println("rpc server: read header error:", err)
		}
		return nil, err
	}
	return &h, nil
}

// 读请求，包括读header
func (server *Server) readRequest(cc codec.Codec) (*request, error) {
	h, err := server.readRequestHeader(cc)
	if err != nil {
		return nil, err
	}

	req := &request{h: h}

	// 通过header 的service method 获取该请求对应service
	req.svc, req.mtype, err = server.findService(h.ServiceMethod)
	if err != nil {
		return req, err
	}

	// 创建对应入参实例
	req.argv = req.mtype.newArgv()
	req.replyv = req.mtype.newReplyv()

	argvi := req.argv.Interface()
	if req.argv.Type().Kind() != reflect.Ptr {
		argvi = req.argv.Addr().Interface()
	}

	if err = cc.ReadBody(argvi); err != nil {
		log.Println("rpc server: read body err:", err)
		return req, err
	}

	return req, nil
}

// 加锁发送消息
func (server *Server) sendResponse(cc codec.Codec, h *codec.Header, body interface{}, sending *sync.Mutex) {
	sending.Lock()
	defer sending.Unlock()
	if err := cc.Write(h, body); err != nil {
		log.Println("rpc server: write response error:", err)
	}
}

// 协程处理单个请求
func (server *Server) handleRequest(cc codec.Codec, req *request, sending *sync.Mutex, wg *sync.WaitGroup, timeout time.Duration) {
	defer wg.Done()

	// 用于接收goroutine的信号, 不用buf chan的原因是在下面超时逻辑需要阻塞掉
	called := make(chan struct{}) // 代表调用成功
	sent := make(chan struct{})   // 代表发送成功

	go func() {
		// 调用请求的服务与方法
		err := req.svc.call(req.mtype, req.argv, req.replyv)
		called <- struct{}{}
		if err != nil {
			req.h.Error = err.Error()
			server.sendResponse(cc, req.h, invalidRequest, sending)
			sent <- struct{}{}
			return
		}
		server.sendResponse(cc, req.h, req.replyv.Interface(), sending)
		sent <- struct{}{}
	}()

	// 超时时间为0，阻塞等待调用结束直接返回
	if timeout == 0 {
		<-called
		<-sent
		return
	}

	// 超时逻辑
	// 此处阻塞获取计时器或者called的值，如果计时器先到达，则直接运行超时逻辑；
	// 如果called先到达，则直接运行sent逻辑发送回应
	// 这样保证了sendResponse只发送一次，要么发送超时逻辑的send，要么发送goroutine里的send
	select {
	case <-time.After(timeout):
		req.h.Error = fmt.Sprintf("rpc server: request handle timeout: expect within %s", timeout)
		server.sendResponse(cc, req.h, invalidRequest, sending)
	case <-called:
		<-sent
	}
}
