package gee_cache

import (
	"fmt"
	"log"
)

const defaultBasePath = "/_geecache/"

type HttpPool struct {
	self     string // 用来记录自己的地址，包括主机名/IP 和端口。
	basePath string // 作为节点间通讯地址的前缀
}

func NewHttpPool(self string) *HttpPool {
	return &HttpPool{
		self:     self,
		basePath: defaultBasePath,
	}
}

func (p *HttpPool) Log(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", p.self, fmt.Sprintf(format, v...))
}
