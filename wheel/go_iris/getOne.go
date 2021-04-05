package main

import (
	"log"
	"net/http"
	"sync"
)

// QPS测试：
// 1. 本机： 4核8线程，16GB内存
// 	wrk -t8 -c500 -d30s 平均QPS在8万多
//
// 2. 服务器A 1核 2GB内存 1Mbps带宽
// 	wrk -t8 -c500 -d30s 平均QPS在700多
//
// 3. 服务器B 16核 64GB内存 5Mbps带宽
// wrk -t8 -c500 -d30s 4000的QPS

var sum int64 = 0

//  预存商品数量
var productNum int64 = 1000000

// 锁
var mutex sync.Mutex

// 计数
var count int64 = 0

// 获取秒杀商品
func GetOneProduct() bool {
	// 加锁
	mutex.Lock()
	defer mutex.Unlock()
	count += 1

	// 判断数据是否超限
	if count%100 == 0 {
		if sum < productNum {
			sum += 1
			log.Println(sum)
			return true
		}
	}

	return false
}

func GetProduct(w http.ResponseWriter, req *http.Request) {
	if GetOneProduct() {
		w.Write([]byte("true"))
		return
	}
	w.Write([]byte("false"))
	return
}

func main() {
	http.HandleFunc("/getOne", GetProduct)
	err := http.ListenAndServe(":8091", nil)
	if err != nil {
		log.Fatal("Err: ", err)
	}
}
