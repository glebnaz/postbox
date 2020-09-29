package mongo

import (
	"github.com/globalsign/mgo"
	"time"
)

//MongoDB is connection struct to mongoDB
//use mgo
type MongoDB struct {
	sess *mgo.Session
}

const connectionTimeout = 10 * time.Second

//NewConnection return new MongoDB connection
//connection with url
//in url must be pass and user if needed
func NewConnection(url string) (*MongoDB, error) {
	var db = MongoDB{}
	var err error
	db.sess, err = mgo.DialWithTimeout(url, connectionTimeout)
	if err != nil {
		return nil, err
	}
	return &db, nil
}

//IsConnected check connection to mongo db Server
func (db *MongoDB) IsConnected() bool {
	return db.sess != nil
}
