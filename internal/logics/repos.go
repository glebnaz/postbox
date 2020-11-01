package logics

import (
	"github.com/glebnaz/postbox/internal/entities"
	"github.com/glebnaz/postbox/internal/errors"
	"github.com/glebnaz/postbox/internal/mongo"
	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
)

//UserRepository interface to storage user Data
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
			return nil, errors.DataBaseOperation.New(err)
		}
	} else {
		q := bson.M{
			"_id": bson.M{
				"$in": ids,
			},
		}
		err := u.db.Find(u.coll, q, &users)
		if err != nil {
			return nil, errors.DataBaseOperation.New(err)
		}
	}
	return users, nil
}

//Insert User to store,return error
func (u UserRepository) Insert(object entities.User) error {
	object.ID = uuid.New().String()
	err := u.db.Insert(u.coll, object)
	if err != nil {
		return errors.DataBaseOperation.New(err)
	}
	return nil
}

//Update user in store by id? remove user with this id,and create new
func (u UserRepository) Update(object entities.User) error {
	if len(object.ID) == 0 {
		return errors.EmptyUsersIDs
	}
	set := bson.M{
		"name":         object.Name,
		"smtp_host":    object.SMTPHost,
		"smtp_address": object.SMTPAddress,
		"smtp_user":    object.SMTPUser,
		"smtp_pass":    object.SMTPPass,
	}
	err := u.db.Update(u.coll, object.ID, set)
	if err != nil {
		return errors.DataBaseOperation.New(err)
	}
	return nil
}

//Delete users by ids,return users and error
func (u UserRepository) Delete(ids []string) (entities.User, error) {
	return entities.User{}, nil
}
