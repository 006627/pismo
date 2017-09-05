package controllers

import (
	"github.com/006627/pismo/transactionsApi/common"
	"github.com/006627/pismo/transactionsApi/models"
	mgo "gopkg.in/mgo.v2"
)

// getContext Information needed to manipulate the database
func getContext(name string) (*mgo.Session, *mgo.Collection) {
	session := common.GetSession().Copy()
	collection := session.DB(common.AppConfig.Database).C(name)

	return session, collection
}

// getNextID increment id in database for not use ObjectId
func getNextID(name string, field string) (int8, error) {
	var id int8
	session, collection := getContext(name)
	defer session.Close()

	if name == "Transactions" {
		result := models.Transaction{}
		err := collection.Find(nil).Sort(field).Limit(1).One(&result)
		if err != nil {
			return -1, err
		}

		id = result.TransactionID + 1
	} else {
		result := models.Register{}
		err := collection.Find(nil).Sort(field).Limit(1).One(&result)
		if err != nil {
			return -1, err
		}

		id = result.PaymentTrackingID + 1
	}

	return id, nil
}
