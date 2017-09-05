package common

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
)

// Session holds the mongodb session for database access
var session *mgo.Session

// GetSession Get database session
func GetSession() *mgo.Session {
	if session == nil {
		createDbSession()
	}
	return session
}

// if the session does not exist a new instance will be created
func createDbSession() {
	var err error
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{AppConfig.MongoDBHost},
		Username: AppConfig.DBUser,
		Password: AppConfig.DBPwd,
		Timeout:  60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[createDbSession]: %s\n", err)
	}
}
