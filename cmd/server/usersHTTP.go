package server

import (
	"github.com/glebnaz/postbox/internal/entities/user"
	locUser "github.com/glebnaz/postbox/internal/logics/user"
	"github.com/glebnaz/postbox/internal/mongo"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserRepository struct {
	coll string
	db   *mongo.MongoDB
}

func initUserRepo(db *mongo.MongoDB) UserRepository {
	return UserRepository{coll: user.Collection, db: db}
}

func (u UserRepository) Get(ids ...string) ([]user.User, error) {
	return nil, nil
}

func (u UserRepository) Insert(object user.User) error {
	return nil
}

func (u UserRepository) Update(object user.User) error {
	return nil
}

func (u UserRepository) Delete(id string) (user.User, error) {
	return user.User{}, nil
}

func (s Server) UserHandler(c echo.Context) error {
	repo := initUserRepo(s.Store)
	if c.Request().Method == http.MethodGet {
		locUser.Get(repo)
	}
	return nil
}
