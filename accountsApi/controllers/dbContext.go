package controllers

import (
	"github.com/006627/pismo/accountsApi/common"
	mgo "gopkg.in/mgo.v2"
)

// GetContext Information needed to manipulate the database
func GetContext(name string) (*mgo.Session, *mgo.Collection) {
	session := common.GetSession().Copy()
	collection := session.DB(common.AppConfig.Database).C(name)

	return session, collection
}
