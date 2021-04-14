package lru

import "container/list"

// lru算法的实现：
// 1. cache与ll底层都是指向Element的指针;
// 2. Element.Value设置为entry
// 3. 每次Get操作将操作的节点移动到ll尾部；
// 4. 每次Add操作判断空间是否超限，如果超限则移除掉ll的头部节点对应的数据；
type Cache struct {
	maxBytes int64                    // 允许使用的最大内存
	nbytes   int64                    // 当前已使用的内存
	ll       *list.List               // Go 语言标准库实现的双向链表list.List。
	cache    map[string]*list.Element // 键是字符串，值是双向链表中对应节点的指针

	OnEvicted func(key string, value Value) // 某条记录被移除时的回调函数
}

// 双向链表节点的数据类型
// 在链表中仍保存每个值对应的 key 的好处在于，淘汰队首节点时，需要用 key 从字典中删除对应的映射。
type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

// New is the Constructor of Cache
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// 查找
// 如果键对应的链表节点存在，则将对应节点移动到队尾，并返回查找到的值。
func (c *Cache) Get(key string) (val Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// 删除
// 实际上是缓存淘汰。即移除最近最少访问的节点（队首）
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back() //  取到队首节点
	if ele != nil {
		c.ll.Remove(ele) // 从链表中删除
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key) // 从字典中删除

		// 更新内存信息
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())

		// 调用回调函数
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// 新增/更新
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok { // 更新
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{
			key:   key,
			value: value,
		})
		c.cache[key] = ele
		c.nbytes += int64(value.Len()) + int64(len(key))
	}

	// 数据大小超过设定的最大值， 调用lru算法移除最少访问的节点
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

// 获取节点个数
func (c *Cache) Len() int {
	return c.ll.Len()
}
