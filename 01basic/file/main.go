package main

import (
	_ "bufio"
	"flag"
	_ "fmt"
)

var (
	ll string
	cat string
)

func init() {
	flag.StringVar(&ll, "ll", "", "show path content")
	flag.StringVar(&cat, "cat", "", "cat file content")
}

/**
go 语言的文件操作  file IO
1. ll 查看某个目录列表
2. cat 查看某个文件内容
*/
func main() {
	flag.Parse()
}


