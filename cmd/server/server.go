package server

import (
	"github.com/glebnaz/postbox/internal/mongo"
	"github.com/labstack/echo/v4"
	"log"
)

const errBadConnectionToStore = "Bad Connection to Store"

//Server main struct of App
//include router as echo
//include Store,ass connection to mongo DB
//Include User and Pass for service
type Server struct {
	router *echo.Echo
	Store  *mongo.MongoDB
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
func InitServer(dbURl string, user, pass string) Server {
	var s Server
	var err error
	s.Store, err = mongo.NewConnection(dbURl)
	if err != nil {
		log.Printf("Error When Connection to Store %v\n", err)
		log.Panic(errBadConnectionToStore)
	}
	return s
}

//Run start server
//use echo Start function
func (s *Server) Run(port string) {
	s.initRouter()
	s.router.Start(port)
}
