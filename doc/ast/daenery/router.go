package main

type Router struct {
	Items []*RouterItem
}

type RouterGroup struct {
	Name   string
	Prefix string
}

type RouterItem struct {
	Method   HttpserverRequestType
	Uri      string
	FuncName string
	Group    *RouterGroup
}

type HttpserverRequestType string

// HttpserverRequestTypes @see git.inke.cn/inkelogic/daenerys/http/server/router.go:10
const (
	Group   HttpserverRequestType = "GROUP"
	Any     HttpserverRequestType = "ANY"
	Get     HttpserverRequestType = "GET"
	Post    HttpserverRequestType = "POST"
	GetPost HttpserverRequestType = "GETPOST"
	// todo 补充
)

func ParseRouterItem() *RouterItem {
	return nil
}
