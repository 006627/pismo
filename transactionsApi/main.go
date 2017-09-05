package main

import (
	"log"
	"net/http"

	"github.com/006627/pismo/transactionsApi/common"
	"github.com/006627/pismo/transactionsApi/routers"
)

func main() {

	router := routers.InitRoutes()
	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: router,
	}

	log.Printf("Listening in Ip %#v\n", common.AppConfig.Server)

	server.ListenAndServe()
}
