package minHeap

type minHeap struct {
	data     []int // 堆数组
	count    int   // 堆的数据大小
	capacity int   // 堆的容量
}

// shiftDown 从下往上整理堆，使其符合最小堆定义
// 约束条件: 父节点必须小于等于子节点，不然交换两个节点
func (m *minHeap) shiftUp(k int) {
	for {
		// 父节点小于子节点，直接break，因为对于shiftUp操作来说
		// 仅仅是为 最后插入的不符合条件的元素k 做的数据重新排列，
		// 如果元素k已经符合条件了，就可以直接break
		if k <= 1 || m.data[k/2] <= m.data[k] {
			break
		}
		m.data[k/2], m.data[k] = m.data[k], m.data[k/2]
		k /= 2
	}
}

// shiftDown 从顶往下整理堆，使其符合最小堆定义
func (m *minHeap) shiftDown(k int) {
	for {
		j := 2 * k
		if j > m.count {
			break
		}

		// 找到两个子节点中的最小值
		if j+1 <= m.count && m.data[j+1] < m.data[j] {
			j++
		}

		// 如果父节点小于等于子节点，直接跳出；否则交换两个节点的值；然后从子节点开始继续往下shiftDown
		if m.data[k] <= m.data[j] {
			break
		}
		m.data[k], m.data[j] = m.data[j], m.data[k]
		k = j
	}
}

// 创建一个空的最小堆
func NewMinHeap(capacity int) *minHeap {
	return &minHeap{
		count:    0,
		data:     make([]int, capacity+1), // 因为最小堆是从1开始计数的，所以实际容量还需要加上0的位置
		capacity: capacity,
	}
}

// 根据数组arr创建一个最小堆
func CreateFromArr(arr []int, n int, capacity int) *minHeap {
	if n > capacity {
		panic("数组长度不能大于堆的容量")
	}

	data := make([]int, capacity+1)
	for i := 0; i < n; i++ {
		data[i+1] = arr[i] // 堆从1开始计数
	}

	m := &minHeap{
		data:     data,
		count:    n,
		capacity: capacity,
	}

	// shiftDown
	for i := m.count / 2; i >= 1; i-- {
		m.shiftDown(i)
	}

	return m
}

func (m *minHeap) Size() int {
	return m.count
}

func (m *minHeap) IsEmpty() bool {
	return m.count == 0
}

func (m *minHeap) Insert(item int) {
	if m.count+1 > m.capacity {
		panic("超出了堆的容量")
	}

	m.data[m.count+1] = item
	m.count++
	m.shiftUp(m.count)
}

func (m *minHeap) ExtractMin() int {
	if m.count <= 0 {
		panic("堆内已经没有数据了")
	}
	ret := m.data[1]

	// 交换头尾节点的位置，然后执行shiftDown
	m.data[1], m.data[m.count] = m.data[m.count], m.data[1]
	m.count--
	m.shiftDown(1)

	return ret
}

func (m *minHeap) GetMin() int {
	if m.count <= 0 {
		panic("堆内已经没有数据了")
	}
	return m.data[1]
}
