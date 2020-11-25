package queue

// 队列的基本应用： 广度优先遍历
// 树：层次遍历
// 图：无权图的最短路径

// 队列通常用双向链表来实现
type (
	// 基于双向链表的队列
	Queue struct {
		top    *node // 队列顶端的结点
		rear   *node // 队列底端的结点
		length int   // 当前队列长度
	}
	// 链表节点
	node struct {
		pre   *node       // 指向前面的结点
		next  *node       // 指向后面的结点
		value interface{} // 结点的值
	}
)

// 获取队列当前长度
func (q *Queue) Len() int {
	return q.length
}

// 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return q.length <= 0
}

// 返回队列顶端元素
func (q *Queue) Peek() interface{} {
	if q.top == nil {
		return nil
	}
	return q.top.value
}

// 入队操作
func (q *Queue) Push(v interface{}) {
	n := &node{
		pre:   nil,
		next:  nil,
		value: v,
	}

	if q.length == 0 { // 队列为空时直接写在top上
		q.top = n
		q.rear = q.top
	} else { // 队列不为空时写在rear上，同时绑定之前rear的结点与新插入结点的关系
		n.pre = q.rear
		q.rear.next = n
		q.rear = n
	}
	q.length++
}

// 推出队列顶端的元素
func (q *Queue) Pop() interface{} {
	if q.length == 0 {
		return nil
	}

	n := q.top
	if q.top.next == nil { // 已经没有元素了
		q.top = nil
	} else {
		q.top = q.top.next
		q.top.pre = nil
	}
	q.length--
	return n.value
}
