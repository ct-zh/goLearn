package main

import (
	"context"
	hot "hot/proto"
	"log"

	"github.com/micro/go-micro/v2"
)

func main() {
	srv := getService()

	acc := "test1"

	res, err := srv.GetAccountInfo(context.TODO(), &hot.AccountRequest{Account: acc})
	if err != nil {
		panic(err)
	}

	log.Printf("get success: %+v", res)
	log.Printf("type %+v %T, %v", res.ClientType, res.ClientType, res.ClientType == hot.Type_Wx)

	//var t hot.Type
	//switch res.ClientType {
	//case hot.Type_App:
	//	t = hot.Type_Wx
	//case hot.Type_Wx:
	//	t = hot.Type_App
	//}
	//
	//log.Printf("set update: type: %v account: %v", res.ClientType, res.Account)
	//_, err = srv.SetAccountType(context.TODO(), &hot.AccountTypeRequest{
	//	Type:    t,
	//	Account: acc,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("update success")
}

func getService() hot.AccountService {
	client := micro.NewService(micro.Name("micro.account.client"), micro.Version("latest"))
	client.Init()

	return hot.NewAccountService("micro.account", client.Client())
}
