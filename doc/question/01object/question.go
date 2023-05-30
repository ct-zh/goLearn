package object

import "fmt"

// question1.

// 定义结构体类型: Teacher, 以及两个方法：getName() 和 getSubject()
// 其中getName方法嵌套了getSubject

type Teacher struct{}

func (t *Teacher) getName() {
	fmt.Println("teacher")
	t.getSubject()
}

func (t *Teacher) getSubject() {
	fmt.Println("math")
}

// 定义结构体类型: Ma, 继承了Teacher（同时继承Teacher的所有方法），重写了getSubject方法

type Ma struct {
	Teacher
}

func (m Ma) getSubject() {
	fmt.Println("wu hu fly")
}
