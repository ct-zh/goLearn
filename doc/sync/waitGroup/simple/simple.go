package main

import (
	"fmt"
	"sync"
)

// 06 一个简单的多协程程序，
// input负责传递「用户输入的数据」到channel
// output负责输出数据

// 知识点： sync.WaitGroup 用来同步
var wg sync.WaitGroup

func main() {
	ch := make(chan string)
	wg.Add(2)
	go input(ch)
	go output(ch)
	wg.Wait()
	fmt.Println("Exit")
}

func input(ch chan string) {
	defer wg.Done()
	defer close(ch)
	var input string
	fmt.Println("input EOF to shut down! ")
	for {
		// 获取输入内容
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			fmt.Println("error: ", err.Error())
			break
		}
		if input == "EOF" {
			fmt.Println("Bye")
			break
		}
		ch <- input
	}
}

func output(ch chan string) {
	defer wg.Done()
	for value := range ch {
		fmt.Println("input: ", value)
	}
}
