package controllers

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/006627/pismo/transactionsApi/common"
)

// Get endpoint to run 'Get' and get all database transactions
func Get(w http.ResponseWriter, r *http.Request) {

	resp, err := findAllTransactions()

	if err != nil {
		displayError(w, err, "Error while try getting all transactions", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// Create end point to execute 'POST' and create a transaction
func Create(w http.ResponseWriter, r *http.Request) {

	resp, err := createSingleTransaction(r.Body)

	if err != nil {
		displayError(w, err, "Error while try create a transaction", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}

// CreatePayments end point to run 'POST' and create payment transactions
func CreatePayments(w http.ResponseWriter, r *http.Request) {
	resp, err := createMultipleTransactions(r.Body)

	if err != nil {
		displayError(w, err, "Error while try getting all transactions", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

//interface to send request to the other micro service
func sendPurchase(body []byte, id int8) error {
	url := fmt.Sprintf("http://%s:8080/v1/accounts/%d", common.AppConfig.Host, id)
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		err = errors.New("Problem whit account limits")
	}

	return err
}

// error handling
func displayError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := []byte(fmt.Sprintf("{\"Error\": %#v, \"Message\": \"%s\", \"HttpStatus\": %d}", handlerError.Error(), message, code))

	log.Printf("[Error]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(errObj)
}
