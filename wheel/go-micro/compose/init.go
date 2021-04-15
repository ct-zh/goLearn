package main

import (
	"encoding/json"
	"github.com/hashicorp/consul/api"
	"log"

	microconf "github.com/LannisterAlwaysPaysHisDebts/goLearn/wheel/go-micro/conf"
)

// 初始化consul配置

const ConsulPath = "127.0.0.1:8500"

func main() {
	conf := config()

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// 设置前缀
	prefix := "hot/"

	kv := client.KV()
	for key, value := range conf {
		p := &api.KVPair{
			Key:   prefix + key,
			Value: value,
		}
		_, err := kv.Put(p, nil)
		if err != nil {
			panic(err)
		}
	}

	log.Println("load success!")
}

func config() map[string][]byte {
	conf := make(map[string][]byte)

	m, err := json.Marshal(microconf.Mysql{Dsn: "root:root@tcp(127.0.0.1:3306)/testdb?charset=utf8&parseTime=True&loc=Local"})
	if err != nil {
		panic(err)
	}
	conf["mysql"] = m

	r, err := json.Marshal(microconf.Redis{
		Host:     "127.0.0.1",
		Port:     "6379",
		Password: "",
	})
	if err != nil {
		panic(err)
	}
	conf["redis"] = r

	return conf
}
