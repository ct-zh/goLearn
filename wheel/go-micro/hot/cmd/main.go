package main

import (
	"hot/common"
	"hot/domain/repository"
	"hot/domain/service"
	"hot/handler"
	hot "hot/proto"
	"log"

	"github.com/micro/go-micro/v2"
)

func main() {
	srv := micro.NewService(
		micro.Name("micro.account"),
		micro.Version("latest"),
		micro.Address(""),
	)

	db, err := common.GetTestDb()
	if err != nil {
		log.Fatal(err)
	}

	accountSrv := service.NewAccount(repository.NewAccount(db))

	err = hot.RegisterAccountHandler(srv.Server(), &handler.Account{Srv: accountSrv})
	if err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
