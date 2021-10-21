package main

import "fmt"

type Teacher struct{}

func (t *Teacher) getName() {
	fmt.Println("teacher")
	t.getSubject()
}

func (t *Teacher) getSubject() {
	fmt.Println("math")
}

type Ma struct {
	Teacher
}

func (m Ma) getSubject() {
	fmt.Println("wu hu fly")
}

func main() {
	m := &Ma{}
	m.getName() // wu hu fly
}
