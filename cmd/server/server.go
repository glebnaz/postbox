package server

import (
	"github.com/glebnaz/postbox/internal/errors"
	"log"

	"github.com/glebnaz/postbox/internal/mongo"
	"github.com/labstack/echo/v4"
)

//Server main struct of App
//include router as echo
//include Store,as connection to mongo DB
//include User and Pass for service
type Server struct {
	router *echo.Echo
	Store  *mongo.DB
	user   string
	pass   string
}

//GetCred return user and pass
//this function need to authorization users
func (s Server) GetCred() (string, string) {
	return s.user, s.pass
}

//InitServer init new Server,return Server
//user pass is field from config
//user and pass is a cred to access server
func InitServer(dbURL string, user, pass string) Server {
	var s Server

	//init cred
	s.user = user
	s.pass = pass

	var err error
	s.Store, err = mongo.NewConnection(dbURL)
	if err != nil {
		log.Panic(errors.DataBaseConnection.New(err))
	}
	return s
}

//Run start server
//use echo Start function
func (s *Server) Run(port string) {
	s.initRouter()
	err := s.router.Start(port)
	if err != nil {
		panic(err)
	}
}
