package stack

import "fmt"

// 栈
type Stack struct {
	capacity int
	top      int
	data     []interface{}
}

func NewStack(capacity int) *Stack {
	return &Stack{
		capacity: capacity,
		top:      -1,
		data:     make([]interface{}, capacity),
	}
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) IsFull() bool {
	return s.top == s.capacity-1
}

func (s *Stack) Push(i interface{}) {
	if s.IsFull() {
		panic("栈已经满了")
	}
	s.top++
	s.data[s.top] = i
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		panic("栈是空的")
	}
	data := s.data[s.top]
	s.top--
	return data
}

func (s *Stack) GetLen() int {
	return s.top + 1
}

func (s *Stack) Clear() {
	s.top = -1
}

func (s *Stack) Traverse() {
	if s.IsEmpty() {
		panic("栈是空的")
	}
	for i := 0; i <= s.top; i++ {
		fmt.Println(s.data[i])
	}
}
