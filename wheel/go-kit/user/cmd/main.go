package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/LannisterAlwaysPaysHisDebts/goLearn/wheel/go-kit/user/common"
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/wheel/go-kit/user/conf"
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/wheel/go-kit/user/dao"
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/wheel/go-kit/user/endpoint"
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/wheel/go-kit/user/service"
	"github.com/LannisterAlwaysPaysHisDebts/goLearn/wheel/go-kit/user/transport"
)

func main() {
	var (
		// 服务器地址与服务器名
		//serviceName = flag.String("service.name", "user", "set service name")
		//serviceHost = flag.String("serivce.host", "127.0.0.1", "set service host")

		// 服务器端口
		servicePort = flag.Int("service.port", 10086, "service port")
	)
	flag.Parse()

	errChan := make(chan error)

	ctx := context.Background()

	// 初始化数据库
	err := dao.InitMysql(conf.InitLocalDb())
	if err != nil {
		log.Fatal(err)
	}

	// 初始化redis
	redisConf := conf.InitLocalRedis()
	err = common.InitRedis(redisConf.Host, redisConf.Port, redisConf.Passwd)
	if err != nil {
		log.Fatal(err)
	}

	// 初始化userService
	userService := service.MakeUserServiceImpl(&dao.UserDAOImpl{})

	userEndpoints := &endpoint.UserEndpoints{
		RegisterEndpoint: endpoint.MakeRegisterEndpoint(userService),
		LoginEndpoint:    endpoint.MakeLoginEndpoint(userService),
	}

	r := transport.MakeHttpHandler(ctx, userEndpoints)

	go func() {
		// 监听http
		errChan <- http.ListenAndServe(":"+strconv.Itoa(*servicePort), r)
	}()

	go func() {
		// 监控系统信号，等待 ctrl + c 系统信号通知服务关闭
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	e := <-errChan
	log.Println(e)
}
