package queue

import (
	"errors"
)

// 优先队列
// 存储数据的 id 与排序值 sort，根据sort降序排列
// 没有参考网上的代码(网上一些代码太水了)，自己写的，可能有bug,谨慎使用
// 使用索引堆来实现

// 创建一个空的优先队列
func CreatePriorityQueue() *priorityQueue {
	return &priorityQueue{
		indexes: []int{0}, // 因为索引从1开始，默认填充一个0
		data:    make(map[int]int),
		reverse: make(map[int]int),
		length:  0,
	}
}

// 获取队列元素数量
func (p *priorityQueue) Len() int {
	return p.length
}

// 判断队列是否为空
func (p *priorityQueue) IsEmpty() bool {
	return p.length == 0
}

// 获取最大的sort与其对应的id
func (p *priorityQueue) GetMax() (id int, sort int) {
	return p.indexes[1], p.data[p.indexes[1]]
}

// 根据id获取其sort
func (p *priorityQueue) GetSort(id int) int {
	return p.data[id]
}

// 判断id是否在队列里面
func (p *priorityQueue) Contain(id int) bool {
	_, ok := p.data[id]
	return ok
}

// 修改id的sort值
func (p *priorityQueue) Change(id int, sort int) error {
	if _, ok := p.data[id]; !ok {
		return errors.New("没有这个数据")
	}

	p.data[id] = sort

	// 向上向下重新排序
	p.shiftUp(p.reverse[id])
	p.shiftDown(p.reverse[id])
	p.dataPrint()
	return nil
}

// 写入一条key-value新数据，根据sort的值进行降序排序
func (p *priorityQueue) Insert(id int, sort int) error {
	l := len(p.indexes)
	if p.length > l-1 { // 因为indexes数组索引从1开始，默认填充了一个0，要比长度多一位
		return errors.New("length字段出现bug，与数组长度不符合")
	} else if p.length == l-1 {
		// todo: resize
		p.indexes = append(p.indexes, id)
	} else {
		p.indexes[p.length+1] = id
	}

	p.data[id] = sort
	p.reverse[id] = p.length + 1
	p.length++
	p.shiftUp(p.length)
	return nil
}

// 弹出sort最大的数据
func (p *priorityQueue) Extract() (id int, sort int, err error) {
	if p.length == 0 {
		return 0, 0, errors.New("当前没有数据")
	}

	id = p.indexes[1]
	sort = p.data[p.indexes[1]]

	// 清除数据
	delete(p.reverse, id)
	delete(p.data, id)
	p.indexes[1] = 0

	p.indexes[1], p.indexes[p.length] = p.indexes[p.length], p.indexes[1]
	p.length--
	p.shiftDown(1)

	return
}
