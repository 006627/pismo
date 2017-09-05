package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Get endpoint to run 'Get' and get all database accounts
func Get(w http.ResponseWriter, r *http.Request) {

	resp, err := GetAccounts()

	if err != nil {
		displayError(w, err, "Error while try getting all accounts", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// Path endpoint to run 'Patch' and make required changes
func Patch(w http.ResponseWriter, r *http.Request) {

	// Get ID from URL Path
	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 8, 0)

	resp, err := PatchAccounts(r.Body, id)

	if err != nil {
		displayError(w, err, "Error while try path all transactions", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

// error handling
func displayError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := []byte(fmt.Sprintf("{\"Error\": %#v, \"Message\": \"%s\", \"HttpStatus\": %d}", handlerError.Error(), message, code))

	log.Printf("[Error]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(errObj)
}
