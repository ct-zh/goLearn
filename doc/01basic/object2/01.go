package main

import (
	"fmt"
	"time"
)

// 字节跳动，高级研发工程师 一面面试题
func main() {
	s := "a"
	for {
		switch s {
		case "a":
			{
				fmt.Println("s1")
				s = "b"
				time.Sleep(time.Second * 1)
				continue
			}
		case "b":
			{
				fmt.Println("s2")
				time.Sleep(time.Second * 1)
				s = "c"
				continue
			}
		case "c":
			{
				fmt.Println("s3")
				time.Sleep(time.Second * 1)
				break // <=== 这个break会跳出循环吗？
			}
		}
	}
}
