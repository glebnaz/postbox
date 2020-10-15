package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func (s *Server) initRouter() {
	s.router = echo.New()
	s.router.Debug = false
	s.router.Logger.SetLevel(log.OFF)

	//get tokens endpoint
	s.router.GET("/token", s.getToken)

	//configure groups
	s.configureEmailGroup()
	s.configureUserGroup()
}

func (s *Server) configureUserGroup() {
	s.router.GET("users", func(c echo.Context) error {
		return c.JSON(200, "OK")
	}, s.middleWare)
}

func (s *Server) configureEmailGroup() {
	emailGroup := s.router.Group("/email")
	emailGroup.Use(s.middleWare)
}
