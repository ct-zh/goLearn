package main

import (
	"flag"
	"fmt"
)

func main() {
	// 将一个参数绑定在变量上面（*变量类型是值*）
	// 参数名为host， 默认值为127.0.0.1, 使用-h参数时帮助文案为: set server ip...
	// 同理还存在BoolVar IntVar 等等 以及interface版的 Var 方法
	var servIp string
	flag.StringVar(&servIp, "host", "127.0.0.1", "set server ip; default 127.0.0.1")

	// 将一个参数的值的指针赋给一个变量（*变量类型是指针*）
	// name: 参数名； value： 默认值； usage： 使用-h 参数时展示的帮助文案
	servPort := flag.Int("port", 10086, "set server port; default 10086")

	servName := flag.String("serv.name", "default service", "set server name")

	// 在所有flag都注册之后，调用Parse 来解析命令行参数写入注册的flag里。
	flag.Parse()

	fmt.Printf(
		"server name: %s \n"+
			"server Ip: %s \n"+
			"server Port: %d \n", *servName, servIp, *servPort)
}
