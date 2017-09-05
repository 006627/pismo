package routers

import (
	"github.com/006627/pismo/transactionsApi/controllers"
	"github.com/gorilla/mux"
)

// api version 1
func addV1Routes(router *mux.Router) {
	router.HandleFunc("/transactions", controllers.Get).Methods("GET")
	router.HandleFunc("/transactions", controllers.Create).Methods("POST")
	router.HandleFunc("/payments", controllers.CreatePayments).Methods("POST")
}

// InitRoutes TODO
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	// v1
	addV1Routes(router.PathPrefix("/v1").Subrouter())

	return router
}
