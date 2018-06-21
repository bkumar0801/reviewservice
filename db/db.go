package db

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
)

/*
NewDBSession ...
*/
func NewDBSession(dbHost, dbUser, dbPassword string) *mgo.Session {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{dbHost},
		Username: dbUser,
		Password: dbPassword,
		Timeout:  60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[createDbSession]: %s\n", err)
	}
	return session
}
