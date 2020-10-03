package mongo

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)

//MongoDB is connection struct to mongoDB
//use mgo
type MongoDB struct {
	sess *mgo.Session
}

const (
	connectionTimeout = 10 * time.Second
)

const errIsNotConnected = "Mongo is not Connected"

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

//Insert insert document to collection
//if collection is not created this function create collection
func (db *MongoDB) Insert(coll string, v ...interface{}) error {
	if !db.IsConnected() {
		return fmt.Errorf("%s", errIsNotConnected)
	}
	var sess = db.sess.Copy()
	defer sess.Close()

	return sess.DB("").C(coll).Insert(v...)
}

//Find find document in collection
func (db *MongoDB) Find(coll string, query map[string]interface{}, v interface{}) error {
	if !db.IsConnected() {
		return fmt.Errorf("%s", errIsNotConnected)
	}
	var sess = db.sess.Copy()
	defer sess.Close()

	bsonQuery := bson.M{}

	for k, qv := range query {
		bsonQuery[k] = qv
	}

	return sess.DB("").C(coll).Find(bsonQuery).All(v)
}

//FindByID find document by ID
func (db *MongoDB) FindByID(coll string, id string, v interface{}) bool {
	if !db.IsConnected() {
		return false
	}
	var sess = db.sess.Copy()
	defer sess.Close()

	return mgo.ErrNotFound != sess.DB("").C(coll).FindId(id).One(v)
}

//FindAll find all document in collection
func (db *MongoDB) FindAll(coll string, v interface{}) error {
	if !db.IsConnected() {
		return fmt.Errorf("%s", errIsNotConnected)
	}
	var sess = db.sess.Copy()
	defer sess.Close()

	return sess.DB("").C(coll).Find(bson.M{}).All(v)
}

//FindWithQuery you can call this function with query
//you can must use mgo.bson format
func (db *MongoDB) FindWithQuery(coll string, query interface{}, v interface{}) error {
	if !db.IsConnected() {
		return fmt.Errorf("%s", errIsNotConnected)
	}
	var sess = db.sess.Copy()
	defer sess.Close()

	return sess.DB("").C(coll).Find(query).One(v)
}

//FindWithQueryAll you can find all document in collection with this function
//you can call this function with mgo.bson query
func (db *MongoDB) FindWithQueryAll(coll string, query interface{}, v interface{}) error {
	if !db.IsConnected() {
		return fmt.Errorf("%s", errIsNotConnected)
	}
	var sess = db.sess.Copy()
	defer sess.Close()

	return sess.DB("").C(coll).Find(query).All(v)
}

//RemoveWithIDs delete all document in collection by ids
func (db *MongoDB) RemoveWithIDs(coll string, ids interface{}) error {
	if !db.IsConnected() {
		return fmt.Errorf("%s", errIsNotConnected)
	}
	var sess = db.sess.Copy()
	defer sess.Close()

	_, err := sess.DB("").C(coll).RemoveAll(bson.M{"_id": bson.M{"$in": ids}})

	return err
}

//Update document by query
//warning you can update all document with this query
func (db *MongoDB) Update(coll string, query interface{}, set interface{}) error {
	if !db.IsConnected() {
		return fmt.Errorf("%s", errIsNotConnected)
	}
	var err error
	var sess = db.sess.Copy()
	defer sess.Close()

	_, err = sess.DB("").C(coll).UpdateAll(query, set)

	return err
}
