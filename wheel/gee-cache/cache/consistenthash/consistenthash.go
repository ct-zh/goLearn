package consistenthash

// 一致性hash实现

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// 函数类型hash
type Hash func(data []byte) uint32

type Map struct {
	hash     Hash           // hash函数，默认crc32
	replicas int            // 虚拟节点倍数，也就是一个节点能有几个虚拟节点
	keys     []int          // hash环
	hashMap  map[int]string // 虚拟节点与真实节点的映射信息,key是节点hash，value是节点值
}

// @replicas 虚拟节点倍数
// @fn hash函数，默认crc32
func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// 添加节点
// @keys 节点列表
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		// 根据虚拟节点倍数创建N个节点
		for i := 0; i < m.replicas; i++ {
			// 虚拟节点的名称为: `strconv.Itoa(i) + key`
			//log.Printf("key:%v add: %v", key, strconv.Itoa(i)+key)

			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash) // 加入hash环
			m.hashMap[hash] = key         // 节点hash映射key
		}
	}
	//log.Println(m.keys)
	sort.Ints(m.keys)
	//log.Println(m.keys)
}

func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	// 计算hash值
	hash := int(m.hash([]byte(key)))

	// 顺时针找到第一个匹配（大于等于key）的虚拟节点 在slice里的下标 idx
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	//log.Printf("key:%v idx:%v", key, idx)

	// 如果 idx == len(m.keys)，说明应选择 m.keys[0]，
	// 因为 m.keys 是一个环状结构，所以用取余数的方式来处理这种情况。
	return m.hashMap[m.keys[idx%len(m.keys)]]
}
