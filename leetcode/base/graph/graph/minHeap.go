package graph

// 用于lazyPrim的最小堆
type minHeap struct {
	data     []Edge // 从1开始的二叉树，存储的内容是edge
	count    int    // 堆中节点的数量
	capacity int    // 堆的容量
}

func (m *minHeap) getVal(k int) float64 {
	return m.data[k].Wt().(float64)
}

func (m *minHeap) shiftUp(i int) {
	for {
		if i <= 1 {
			break
		}
		if m.getVal(i/2) <= m.getVal(i) {
			break
		}
		m.data[i/2], m.data[i] = m.data[i], m.data[i/2]
		i /= 2
	}
}

func (m *minHeap) shiftDown(i int) {
	for {
		if i*2 > m.count { // 判断是否存在子节点
			break
		}
		j := i * 2
		if j+1 <= m.count && m.getVal(j+1) < m.getVal(j) {
			j = j + 1
		}
		if m.getVal(i) <= m.getVal(j) {
			break
		}
		m.data[i], m.data[j] = m.data[j], m.data[i]
		i = j
	}
}

func NewMinHeap(capacity int) *minHeap {
	m := &minHeap{
		data:     make([]Edge, capacity+1), // 从1开始的数组
		count:    0,
		capacity: capacity,
	}
	m.data[0] = Edge{}

	return m
}

func (m *minHeap) Count() int {
	return m.count
}

func (m *minHeap) IsEmpty() bool {
	return m.count == 0
}

func (m *minHeap) Insert(i Edge) {
	if m.count >= m.capacity {
		panic("该堆容量已满")
	}
	m.data[m.count+1] = i
	m.count++
	m.shiftUp(m.count)
}

// 推出最小值
func (m *minHeap) ExtractMin() Edge {
	if m.count == 0 {
		panic("是个空堆")
	}
	ref := m.data[1]
	m.data[1], m.data[m.count] = m.data[m.count], m.data[1]
	m.count--
	m.shiftDown(1)
	return ref
}
