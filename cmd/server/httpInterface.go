package server

import "github.com/labstack/echo"

type response struct {
	Status string      `json:"status,omitempty"`
	Error  string      `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func (s *Server) initRouter() {
	s.router = echo.New()
	//configure groups

	s.configureEmailGroup()
	s.configureUserGroup()
}

func (s *Server) configureUserGroup() {
	userGroup := s.router.Group("/users")
	userGroup.Use(s.middleWare)

}

func (s *Server) configureEmailGroup() {
	emailGroup := s.router.Group("/email")
	emailGroup.Use(s.middleWare)
}
