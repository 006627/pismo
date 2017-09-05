package main

import (
	"log"
	"net/http"

	"github.com/006627/pismo/accountsApi/common"
	"github.com/006627/pismo/accountsApi/routers"
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
