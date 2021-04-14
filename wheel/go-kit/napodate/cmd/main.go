package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/LannisterAlwaysPaysHisDebts/goLearn/wheel/go-kit/napodate"
)

func main() {
	var (
		httpAddr = flag.String("http", ":12300", "http listen address")
	)
	flag.Parse()

	ctx := context.Background()

	srv := napodate.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// 映射服务
	endpoints := napodate.Endpoints{
		GetEndPoint:      napodate.MakeGetEndpoint(srv),
		StatusEndPoint:   napodate.MakeStatusEndpoint(srv),
		ValidateEndPoint: napodate.MakeValidateEndpoint(srv),
	}

	// http传输
	go func() {
		log.Println("napodate is listening on port: ", *httpAddr)
		handler := napodate.NewHttpServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
