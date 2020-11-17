package tree

import "fmt"

// 顺序查找表 sequence search Tree
// 使用链表实现，新节点直接插在链表的头部

type SST struct {
	head  *sstNode
	count int
}

type sstNode struct {
	key   int
	value string
	next  *sstNode
}

func (s *SST) Size() int {
	return s.count
}

func (s *SST) IsEmpty() bool {
	return s.count == 0
}

// 向顺序查找表中插入一个新的(key, value)数据对
func (s *SST) Insert(key int, value string) {
	// 查找一下整个顺序表，是否存在同样大小的key
	cur := s.head
	for cur != nil {
		// 若在顺序表中找到了同样大小key的节点
		// 则当前节点不需要插入，将该key所对应的值更新为value后返回
		if key == cur.key {
			cur.value = value
			return
		}
		cur = cur.next
	}

	// 若顺序表中没有同样大小的key，则创建新节点，将新节点直接插在表头
	newNode := &sstNode{
		key:   key,
		value: value,
	}
	newNode.next, s.head = s.head, newNode
	s.count++
}

func (s *SST) Contain(key int) bool {
	cur := s.head
	for cur != nil {
		if cur.key == key {
			return true
		}
		cur = cur.next
	}
	return false
}

func (s *SST) Search(key int) (string, bool) {
	cur := s.head
	for cur != nil {
		if cur.key == key {
			return cur.value, true
		}
		cur = cur.next
	}
	return "", false
}

func (s *SST) Remove(key int) {
	if s.head == nil {
		return
	}

	dummyNode := &sstNode{next: s.head}

	cur := dummyNode
	for cur.next != nil && cur.next.key != key {
		cur = cur.next
	}

	// 说明找到了,cur.next.key = key
	if cur.next != nil {
		delNode := cur.next
		cur.next = delNode.next
		s.count--

		// 防止删除的是头节点
		s.head = dummyNode.next
	}
}

func (s *SST) Print() {
	if s.head == nil {
		return
	}

	cur := s.head
	for cur.next != nil {
		fmt.Printf(" %d:%s -> ", cur.key, cur.value)
		cur = cur.next
	}
	fmt.Printf(" %d:%s -> NULL\n", cur.key, cur.value)
}
