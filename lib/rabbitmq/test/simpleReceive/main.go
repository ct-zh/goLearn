package main

import "github.com/LannisterAlwaysPaysHisDebts/goLearn/lib/rabbitmq"

func main() {
	rabbit := rabbitmq.NewSimple("imoocSimple")
	rabbit.ConsumeSimple()
}
