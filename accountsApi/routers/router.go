package routers

import (
	"github.com/006627/pismo/accountsApi/controllers"
	"github.com/gorilla/mux"
)

// api version 1
func addV1Routes(router *mux.Router) {
	router.HandleFunc("/accounts/limits", controllers.Get).Methods("GET")
	router.HandleFunc("/accounts/{id}", controllers.Patch).Methods("PATCH")
}

// InitRoutes TODO
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	// v1
	addV1Routes(router.PathPrefix("/v1").Subrouter())

	return router
}
