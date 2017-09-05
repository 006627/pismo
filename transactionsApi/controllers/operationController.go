package controllers

import (
	"errors"
	"fmt"

	"github.com/006627/pismo/transactionsApi/models"
	"gopkg.in/mgo.v2/bson"
)

const (
	ePURCHASE     = 1
	eINSTALLMENTS = 2
	eWITHDRAWAL   = 3
	ePAYMENT      = 4
)

// records the payment in the PaymentsTracking table
func registerPayment(creditID int8, debitID models.Transaction, amount float32) error {

	var register models.Register
	id, err := getNextID("PaymentsTracking", "-PaymentTracking_ID")
	if err != nil {
		id = 1
	}
	register.PaymentTrackingID = id
	register.CreditTransactionID = creditID
	register.DebitTransactionID = debitID.TransactionID
	register.Amount = amount

	session, collection := getContext("PaymentsTracking")
	defer session.Close()

	err = collection.Insert(&register)

	debitID.Amount = amount

	//after registration the notification is sent to the other API
	err = purchaseOperation(&debitID, debitID.OperationTypeID)

	return err
}

//filters transactions by Operation id
func executeOperation(t *models.Transaction) error {
	var err error
	switch t.OperationTypeID {
	case ePURCHASE, eINSTALLMENTS, eWITHDRAWAL:
		//amount must be negative
		if t.Amount > 0 {
			t.Amount = t.Amount * (-1)
			t.Balance = t.Balance * (-1)
		}
		err = purchaseOperation(t, t.OperationTypeID)
	case ePAYMENT:
		paymentOperation(t)
	default:
		err = errors.New("Invalid Operation")
	}

	return err
}

//mounts the message to be sent to another API
func purchaseOperation(t *models.Transaction, op int8) error {

	var operation string

	if op == eWITHDRAWAL {
		operation = "available_withdrawal_limit"
	} else {
		operation = "available_credit_limit"
	}
	//according to the specification the endpoint consumes
	//a different structure from the database schema
	body := []byte(fmt.Sprintf("{\"%s\": { \"amount\": %f }}", operation, t.Amount))
	err := sendPurchase(body, t.AccountID)

	return err
}

// manipulates the database to execute the payments
func paymentOperation(t *models.Transaction) error {

	session, collection := getContext("Transactions")
	defer session.Close()

	// filter for the account, the type of operation and for the negative balances
	// orders by ChangeOrder and EventDate  as per specification
	pipeline := []bson.M{
		{"$match": bson.M{"Account_ID": t.AccountID}},
		{"$match": bson.M{"OperationType_ID": bson.M{"$ne": ePAYMENT}}},
		{"$match": bson.M{"Balance": bson.M{"$ne": 0}}},
		{"$lookup": bson.M{
			"from":         "OperationsTypes",
			"localField":   "OperationType_ID",
			"foreignField": "OperationType_ID",
			"as":           "op"}},
		{"$project": bson.M{
			"_id":              0,
			"Transaction_ID":   1,
			"Account_ID":       1,
			"Balance":          1,
			"EventDate":        1,
			"OperationType_ID": 1,
			"charge":           "$op.ChargeOrder"}},
		{"$sort": bson.M{
			"charge":    1,
			"EventDate": 1}}}

	pipe := collection.Pipe(pipeline)

	iter := pipe.Iter()
	resp := models.Transaction{}

	//for each negative balance is performed the registration
	//and the updating in the database
	for iter.Next(&resp) {

		if t.Balance > 0 {
			if t.Balance+resp.Balance <= 0 {
				err := registerPayment(t.TransactionID, resp, t.Balance)
				if err != nil {
					return err
				}
				resp.Balance += t.Balance
				t.Balance = 0
			} else {
				err := registerPayment(t.TransactionID, resp, -resp.Balance)
				if err != nil {
					return err
				}
				t.Balance += resp.Balance
				resp.Balance = 0
			}

			err := updateTransaction(&resp)

			if err != nil {
				return err
			}
		}
	}

	//when the payment amount is written down, it is updated
	err := updateTransaction(t)
	if err != nil {
		return err
	}

	return nil
}
