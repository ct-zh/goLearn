package main

import (
	"fmt"
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/lib/rabbitmq"
)

func main() {
	rabbit := rabbitmq.NewSimple("imoocSimple")
	rabbit.PublishSimple("Hello imooc!")
	fmt.Println("发送成功")
}
