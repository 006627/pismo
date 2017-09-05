package models

import (
	"encoding/json"
	"time"
)

// Transaction Model
type Transaction struct {
	TransactionID   int8      `bson:"Transaction_ID" json:"TransactionID"`
	AccountID       int8      `bson:"Account_ID" json:"account_id"`
	OperationTypeID int8      `bson:"OperationType_ID" json:"operation_type_id"`
	Amount          float32   `bson:"Amount" json:"amount"`
	Balance         float32   `bson:"Balance" json:"balance"`
	EventDate       time.Time `bson:"EventDate" json:"eventDate"`
	DueDate         time.Time `bson:"DueDate" json:"dueDate"`
}

// Register Model
type Register struct {
	PaymentTrackingID   int8    `bson:"PaymentTracking_ID" json:"PaymentTrackingID"`
	CreditTransactionID int8    `bson:"CreditTransaction_ID" json:"CreditTransactionID"`
	DebitTransactionID  int8    `bson:"DebitTransaction_ID" json:"DebitTransactionID"`
	Amount              float32 `bson:"amount" json:"amount"`
}

// Transactions slice Model
type Transactions []Transaction

// ToJSON transform Transaction in JSON
func (t *Transaction) ToJSON() []byte {
	ToJSON, err := json.Marshal(t)

	if err != nil {
		panic(err)
	}

	return ToJSON
}

// ToJSON ransform Transaction in JSON
// TODO create a interface
func (t *Transactions) ToJSON() []byte {
	ToJSON, err := json.Marshal(t)

	if err != nil {
		panic(err)
	}

	return ToJSON
}
