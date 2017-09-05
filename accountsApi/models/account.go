package models

import (
	"encoding/json"
	"errors"

	"gopkg.in/mgo.v2/bson"
)

// Account Model
type Account struct {
	ID                       bson.ObjectId `bson:"_id,omitempty" json:"-"`
	AccountID                int8          `bson:"Account_ID,omitempty" json:"account_ID,omitempty"`
	AvailableCreditLimit     float32       `bson:"AvailableCreditLimit" json:"available_credit_limit"`
	AvailableWithdrawalLimit float32       `bson:"AvailableWithdrawalLimit" json:"available_withdrawal_limit"`
}

// Accounts slice Model
type Accounts []Account

// ToJSON transform slice Account in JSON
func (a *Accounts) ToJSON() []byte {
	ToJSON, err := json.Marshal(a)

	if err != nil {
		panic(err)
	}

	return ToJSON
}

// ToJSON transform Account in JSON
// TODO create a interface
func (a *Account) ToJSON() []byte {
	ToJSON, err := json.Marshal(a)

	if err != nil {
		panic(err)
	}

	return ToJSON
}

// FromJSON transform JSON in account
func FromJSON(data []byte) Account {
	account := Account{}
	err := json.Unmarshal(data, &account)

	if err != nil {
		panic(err)
	}

	return account
}

// UnmarshalJSON Implement custom method
func (a *Account) UnmarshalJSON(data []byte) error {
	var tmp interface{}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	md, ok := tmp.(map[string]interface{})
	if !ok {
		return errors.New("ERROR")
	}

	a.AccountID = 0
	if md["account_ID"] != nil {
		a.AccountID = int8(md["account_ID"].(float64))
	}

	// TODO
	a.AvailableCreditLimit = 0
	if md["available_credit_limit"] != nil {
		aux := md["available_credit_limit"].(map[string]interface{})
		a.AvailableCreditLimit = float32(aux["amount"].(float64))
	}

	a.AvailableWithdrawalLimit = 0
	if md["available_withdrawal_limit"] != nil {
		aux := md["available_withdrawal_limit"].(map[string]interface{})
		a.AvailableWithdrawalLimit = float32(aux["amount"].(float64))
	}

	return nil
}
