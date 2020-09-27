package main

import (
	"fmt"
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/lib/rabbitmq"
	"strconv"
	"time"
)

func main() {
	rabbit := rabbitmq.NewSimple("" + "imoocSimple")
	for i := 0; i <= 100; i++ {
		rabbit.PublishSimple("Hello imooc ! " + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
