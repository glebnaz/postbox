package logics

import (
	"github.com/glebnaz/postbox/internal/entities"
	"github.com/glebnaz/postbox/internal/mongo"
	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
)

type UserRepository struct {
	coll string
	db   *mongo.DB
}

//InitUserRepo create UserRepo
func InitUserRepo(db *mongo.DB) UserRepository {
	return UserRepository{coll: entities.Collection, db: db}
}

//Get return users,find by id,if ids empty find all
func (u UserRepository) Get(ids []string) ([]entities.User, error) {
	var users []entities.User
	if len(ids) == 0 {
		err := u.db.FindAll(u.coll, &users)
		if err != nil {
			return nil, err
		}
	} else {
		q := bson.M{
			"_id": bson.M{
				"$in": ids,
			},
		}
		err := u.db.Find(u.coll, q, &users)
		if err != nil {
			return nil, err
		}
	}
	return users, nil
}

func (u UserRepository) Insert(object entities.User) error {
	object.ID = uuid.New().String()
	return u.db.Insert(u.coll, object)
}

func (u UserRepository) Update(object entities.User) error {
	return nil
}

func (u UserRepository) Delete(ids []string) (entities.User, error) {
	return entities.User{}, nil
}
