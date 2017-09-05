package models_test

import (
	"reflect"
	"testing"

	"github.com/006627/pismo/accountsApi/models"
)

func TestAccountToJSON(t *testing.T) {
	t.Log("--ToJSON--")
	account := models.Account{AccountID: 1, AvailableCreditLimit: 3.14, AvailableWithdrawalLimit: 5.15}
	json := account.ToJSON()

	assert := `{"account_ID":1,"available_credit_limit":3.14,"available_withdrawal_limit":5.15}`
	test := reflect.DeepEqual(string(json), assert)

	if test != true {
		t.Errorf("Expected %s, but it was %s instead.", assert, string(json))
	}

}

func TestAccountUnmarshalJSON(t *testing.T) {
	t.Log("--UnmarshalJSON--")

	var account models.Account
	json := `{
		"available_credit_limit": {
			"amount": -123.45
		},
	
		"available_withdrawal_limit": {
			"amount": -123.45
		}
	}`

	account.UnmarshalJSON([]byte(json))

	if account.AvailableCreditLimit != -123.45 {
		t.Errorf("Expected -123.45, but it was %f instead.", account.AvailableCreditLimit)
	}

	if account.AvailableWithdrawalLimit != -123.45 {
		t.Errorf("Expected -123.45, but it was %f instead.", account.AvailableWithdrawalLimit)
	}
}
