package main

import (
	"hot/common"
	conf2 "hot/conf"
	"hot/domain/repository"
	"hot/domain/service"
	"hot/handler"
	hot "hot/proto"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/config/source/consul/v2"
	consulRegister "github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	// {{{{{{ 注册配置中心
	consulSource := consul.NewSource(
		//设置配置中心的地址
		consul.WithAddress("127.0.0.1:8500"),
		//设置前缀，不设置默认前缀 /micro/config
		consul.WithPrefix("hot/"),
		//是否移除前缀，这里是设置为true，表示可以不带前缀直接获取对应配置
		consul.StripPrefix(true),
	)

	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	if err = conf.Load(consulSource); err != nil {
		log.Fatal(err)
	}
	// }}}}}}}

	// {{{{  注册服务中心
	consulRegister := consulRegister.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})
	// }}}}

	srv := micro.NewService(
		micro.Name("micro.account"),
		micro.Version("latest"),
		micro.Address(""),
		micro.Registry(consulRegister),
	)

	var mysql conf2.Mysql
	err = conf.Get("mysql").Scan(&mysql)
	if err != nil {
		log.Fatal("加载mysql配置失败: ", err)
	}

	db, err := gorm.Open("mysql", mysql.Dsn)
	if err != nil {
		log.Fatal(err)
	}
	db.SingularTable(true)

	var r conf2.Redis
	err = conf.Get("redis").Scan(&r)
	if err != nil {
		panic(err)
	}
	_ = common.InitRedis(r.Host, r.Port, r.Password)
	redisConn, err := common.GetRedisConn()
	if err != nil {
		log.Fatal(err)
	}

	accountSrv := service.NewAccount(repository.NewAccount(db, redisConn))

	err = hot.RegisterAccountHandler(srv.Server(), &handler.Account{Srv: accountSrv})
	if err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
