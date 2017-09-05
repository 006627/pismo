package controllers

import (
	"errors"
	"io"
	"io/ioutil"

	"gopkg.in/mgo.v2/bson"

	"github.com/006627/pismo/accountsApi/models"
)

// GetAccounts handler for GET
func GetAccounts() ([]byte, error) {
	var accounts models.Accounts

	session, collection := GetContext("Accounts")
	defer session.Close()

	err := collection.Find(nil).All(&accounts)

	if err != nil {
		return nil, err
	}

	return accounts.ToJSON(), nil
}

// GetAccountByID needed for the handler PATCH
func getAccountByID(id int64) (models.Account, error) {

	var account models.Account

	session, collection := GetContext("Accounts")
	defer session.Close()

	err := collection.Find(bson.M{"Account_ID": id}).One(&account)

	return account, err
}

// PatchAccounts handler for PATCH
func PatchAccounts(b io.ReadCloser, id int64) ([]byte, error) {

	accUpdate, err := getAccountByID(id)
	if err != nil {
		return nil, err
	}

	// Get Body Data
	body, err := ioutil.ReadAll(b)
	if err != nil {
		return nil, err
	}

	acc := models.FromJSON(body)

	// control to not exceed the limit
	if acc.AvailableCreditLimit < 0 && -acc.AvailableCreditLimit > accUpdate.AvailableCreditLimit {
		return nil, errors.New("The credit limit is over")
	}
	accUpdate.AvailableCreditLimit += acc.AvailableCreditLimit

	if acc.AvailableWithdrawalLimit < 0 && -acc.AvailableWithdrawalLimit > accUpdate.AvailableWithdrawalLimit {
		return nil, errors.New("The Withdrawal limit is over")
	}
	accUpdate.AvailableWithdrawalLimit += acc.AvailableWithdrawalLimit

	session, collection := GetContext("Accounts")
	defer session.Close()

	err = collection.Update(bson.M{"Account_ID": accUpdate.AccountID},
		bson.M{"$set": bson.M{
			"AvailableCreditLimit":     accUpdate.AvailableCreditLimit,
			"AvailableWithdrawalLimit": accUpdate.AvailableWithdrawalLimit,
		}})

	return accUpdate.ToJSON(), err
}
