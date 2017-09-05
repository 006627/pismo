package controllers

import (
	"encoding/json"
	"io"
	"time"

	"github.com/006627/pismo/transactionsApi/models"
	"gopkg.in/mgo.v2/bson"
)

// FindAllTransactions handler for GET
func findAllTransactions() ([]byte, error) {
	var transactions models.Transactions

	session, collection := getContext("Transactions")
	defer session.Close()

	err := collection.Find(nil).All(&transactions)

	if err != nil {
		return nil, err
	}

	return transactions.ToJSON(), nil
}

// CreateSingleTransaction handler for POST (transaction)
func createSingleTransaction(body io.ReadCloser) ([]byte, error) {
	var transaction models.Transaction

	err := json.NewDecoder(body).Decode(&transaction)
	if err != nil {
		return nil, err
	}
	return createTransaction(&transaction)
}

// CreateMultipleTransactions handler for POST (Payments)
func createMultipleTransactions(body io.ReadCloser) ([]byte, error) {
	var transactions models.Transactions

	err := json.NewDecoder(body).Decode(&transactions)
	if err != nil {
		return nil, err
	}

	for _, t := range transactions {
		t.OperationTypeID = ePAYMENT
		createTransaction(&t)
	}

	return transactions.ToJSON(), nil
}

// CreateTransaction handler for POST
// handles the transaction and forwards it to other processes
func createTransaction(t *models.Transaction) ([]byte, error) {
	id, err := getNextID("Transactions", "-Transaction_ID")
	if err != nil {
		id = 1
	}

	t.TransactionID = id
	t.EventDate = time.Now()
	t.DueDate = t.EventDate.AddDate(0, 0, 20)
	t.Balance = t.Amount

	session, collection := getContext("Transactions")
	defer session.Close()

	//filters by Operation id
	err = executeOperation(t)
	if err != nil {
		return nil, err
	}

	err = collection.Insert(t)
	if err != nil {
		return nil, err
	}

	//if there are payments with the positive balance and amounts payable
	//the transaction is carried out
	if t.OperationTypeID != ePAYMENT {
		err = verifyPaymentsCredits(t.AccountID)
	}

	return t.ToJSON(), err

}

// used to update transactions after payments are made
func updateTransaction(t *models.Transaction) error {
	session, collection := getContext("Transactions")
	defer session.Close()

	err := collection.Update(bson.M{"Transaction_ID": t.TransactionID},
		bson.M{"$set": bson.M{"Balance": t.Balance}})

	if err != nil {
		return err
	}

	return nil

}

// responsible for outstanding payments
func verifyPaymentsCredits(id int8) error {
	session, collection := getContext("Transactions")
	defer session.Close()

	// filter by amounts payable
	pipeline := []bson.M{
		{"$match": bson.M{"Account_ID": id}},
		{"$match": bson.M{"OperationType_ID": ePAYMENT}},
		{"$match": bson.M{"Balance": bson.M{"$gt": 0}}}}
	pipe := collection.Pipe(pipeline)

	iter := pipe.Iter()
	resp := models.Transaction{}

	// each pending payment is sent to the payment cycle again
	for iter.Next(&resp) {
		err := paymentOperation(&resp)
		if err != nil {
			return err
		}
	}

	return nil
}
