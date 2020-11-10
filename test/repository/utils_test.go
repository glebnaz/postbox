package repository

import (
	"io/ioutil"
	"os"

	"github.com/glebnaz/postbox/internal/mongo"
)

func initDBSession() *mongo.DB {
	urlDB := os.Getenv("DB_URL_TEST")
	if len(urlDB) == 0 {
		panic("Empty URL DB")
	}
	sess, err := mongo.NewConnection(urlDB)
	if err != nil {
		panic(err)
	}
	return sess
}

func PrepareTest(repo string) (*mongo.DB, []byte) {
	sess := initDBSession()
	switch repo {
	case UserRepoType:
		file, err := os.Open("test/repository/moks/user.json")
		if err != nil {
			panic(err)
		}
		data, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		return sess, data
	}
	return nil, nil
}

const UserRepoType = "user"
