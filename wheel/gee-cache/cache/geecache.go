package cache

import (
	"fmt"
	geecache "geecache/cache/proto"
	"geecache/cache/singleflight"
	"log"
	"sync"
)

type Getter interface {
	Get(key string) ([]byte, error)
}

// GetterFunc是接口型函数，该函数实现了Getter接口
type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

// Group类似命名空间，不同的命名空间对应不同的缓存空间
// 回调函数getter，当数据源不存在时调用getter获取数据;获取数据的方法交给用户
type Group struct {
	name      string
	getter    Getter
	mainCache cache
	peers     PeerPicker

	// 使用singleFlight来做缓冲器
	// 每个key在同一时间只会被请求一次
	// 更详细的实现见 https://golang.org/x/sync/singleflight
	loader *singleflight.Group
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group) // 全局所有命名空间集合
)

// 新建group
// @name 命名空间名称
// @cacheBytes 缓存大小
// @getter
func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil getter")
	}
	mu.Lock()
	defer mu.Unlock()

	// 没有判断is set
	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytes: cacheBytes},
		loader:    &singleflight.Group{},
	}
	groups[name] = g
	return g
}

func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()
	return g
}

// 获取key
func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}

	// 缓存命中，则返回数据
	if v, ok := g.mainCache.get(key); ok {
		log.Println("[GeeCache] hit")
		return v, nil
	}

	// 缓存不存在，则调用load方法
	return g.load(key)
}

// 将 实现了 PeerPicker 接口的 HTTPPool 注入到 Group 中。
func (g *Group) RegisterPeers(peers PeerPicker) {
	if g.peers != nil {
		panic("RegisterPeerPicker called more than once")
	}
	g.peers = peers
}

//
// 分布式场景下会调用 getFromPeer 从其他节点获取
// load 调用 getLocally
func (g *Group) load(key string) (value ByteView, err error) {
	viewi, err := g.loader.Do(key, func() (interface{}, error) {
		if g.peers != nil {
			if peer, ok := g.peers.PickPeer(key); ok {
				if value, err = g.getFromPeer(peer, key); err == nil {
					return value, nil
				}
				log.Println("[GeeCache] Failed to get from peer", err)
			}
		}
		return g.getLocally(key)
	})

	if err == nil {
		return viewi.(ByteView), nil
	}

	return
}

// getLocally 调用用户回调函数 g.getter.Get() 获取源数据
func (g *Group) getLocally(key string) (ByteView, error) {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	value := ByteView{b: cloneBytes(bytes)}
	g.populateCache(key, value)
	return value, nil
}

func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}

func (g *Group) getFromPeer(peer PeerGetter, key string) (ByteView, error) {
	req := &geecache.Request{
		Group: g.name,
		Key:   key,
	}

	res := &geecache.Response{}
	err := peer.Get(req, res)
	if err != nil {
		return ByteView{}, err
	}
	return ByteView{b: res.Value}, nil
}
