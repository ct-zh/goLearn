package main

import (
	"fmt"
	"sync"
)

// waitGroup用法
var wg sync.WaitGroup

func main() {
	ch := make(chan string)
	wg.Add(2) // 总共需要等待两个Done执行
	go input(ch)
	go output(ch)
	wg.Wait() // 阻塞等待两个Done执行完毕
	fmt.Println("Exit")
}

func input(ch chan string) {
	defer wg.Done()
	defer close(ch) // 一对一的 channel 必须在发送端进行关闭操作
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
