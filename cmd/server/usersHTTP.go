package server

import (
	"github.com/glebnaz/postbox/internal/logics"
	"github.com/labstack/echo/v4"
	"net/http"
)

//UserHandler implements a method for returning, updating and deleting users
func (s Server) UserHandler(c echo.Context) error {
	repo := logics.InitUserRepo(s.Store)
	var err error
	if c.Request().Method == http.MethodGet {
		var req logics.UserReq
		err = c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, logics.UserResp{Status: "error", Error: err.Error()})
		}
		users, err := logics.GetUsers(repo, req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, logics.UserResp{Status: "error", Error: err.Error()})
		}
		return c.JSON(http.StatusOK, logics.UserResp{Status: "success", Users: users})
	}
	return c.JSON(http.StatusMethodNotAllowed, logics.UserResp{Status: "error"})
}
