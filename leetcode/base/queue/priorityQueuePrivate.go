package queue

import "fmt"

// 私有变量与私有方法

// 优先队列
type priorityQueue struct {
	indexes []int       // 按照data里的weight大小排序 value: id
	data    map[int]int // key: id; value: weight
	reverse map[int]int // key: id value:indexes的id 在shiftUp与shiftDown时需要
	length  int
}

func (p *priorityQueue) shiftDown(k int) {
	for k >= 1 && k*2 <= p.length {
		j := k * 2
		if j+1 <= p.length && p.data[p.indexes[j+1]] > p.data[p.indexes[j]] {
			j++
		}

		if p.data[p.indexes[k]] > p.data[p.indexes[j]] {
			break
		}
		p.indexes[k], p.indexes[j] = p.indexes[j], p.indexes[k]
		p.reverse[p.indexes[k]] = k
		p.reverse[p.indexes[j]] = j
		k = j
	}
}

func (p *priorityQueue) shiftUp(k int) {
	for k > 1 && k <= p.length &&
		p.data[p.indexes[k]] > p.data[p.indexes[k/2]] {
		p.indexes[k], p.indexes[k/2] = p.indexes[k/2], p.indexes[k]
		p.reverse[p.indexes[k]] = k
		p.reverse[p.indexes[k/2]] = k / 2
		k /= 2
	}
}

func (p *priorityQueue) dataPrint() {
	fmt.Printf("[")
	for _, value := range p.indexes {
		if value != 0 {
			fmt.Printf("%d:%d ", value, p.data[value])
		}
	}
	fmt.Printf("]\n")

}
