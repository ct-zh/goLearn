package cache

import (
	"fmt"
	"geecache/cache/consistenthash"
	geecache "geecache/cache/proto"
	"github.com/golang/protobuf/proto"
	"log"
	"net/http"
	"strings"
	"sync"
)

const (
	defaultBasePath = "/_geecache/"
	defaultReplicas = 50 // 默认50倍虚拟节点
)

type HttpPool struct {
	self     string // 用来记录自己的地址，包括主机名/IP 和端口。
	basePath string // 作为节点间通讯地址的前缀

	mu          sync.Mutex             // 写锁
	peers       *consistenthash.Map    // 一致性hash map
	httpGetters map[string]*httpGetter // 远程访问节点map
}

func NewHttpPool(self string) *HttpPool {
	return &HttpPool{
		self:     self,
		basePath: defaultBasePath,
	}
}

// 启动http服务
func (p *HttpPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 判断url前缀是否是basePath
	if !strings.HasPrefix(r.URL.Path, p.basePath) {
		panic("HttpPool serving unexpected path: " + r.URL.Path)
	}

	p.Log("%s %s", r.Method, r.URL.Path)

	// 约定访问路径格式为： <basepath>/<groupname>/<key> required
	// 地址/命名空间/key
	parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	groupName := parts[0]
	key := parts[1]

	group := GetGroup(groupName)
	if group == nil {
		http.Error(w, "no such group: "+groupName, http.StatusNotFound)
		return
	}

	view, err := group.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := proto.Marshal(&geecache.Response{
		Value: view.ByteSlice(),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(body)
}

// 设置所有节点
// @peers 节点地址
func (p *HttpPool) Set(peers ...string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// 创建一致性hash map
	p.peers = consistenthash.New(defaultReplicas, nil)
	p.peers.Add(peers...)

	// 创建peer string 与对应的httpGetter映射
	p.httpGetters = make(map[string]*httpGetter, len(peers))
	for _, peer := range peers {
		p.httpGetters[peer] = &httpGetter{baseURL: peer + p.basePath}
	}
}

// 选择一个节点
func (p *HttpPool) PickPeer(key string) (PeerGetter, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if peer := p.peers.Get(key); peer != "" && peer != p.self {
		p.Log("Pick peer %s", peer)
		return p.httpGetters[peer], true
	}
	return nil, false
}

// httpPool实现了PeerPicker接口
var _ PeerPicker = (*HttpPool)(nil)

func (p *HttpPool) Log(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", p.self, fmt.Sprintf(format, v...))
}
