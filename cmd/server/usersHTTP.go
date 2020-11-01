package server

import (
	"net/http"

	"github.com/glebnaz/postbox/internal/logics"
	"github.com/labstack/echo/v4"
)

//UserHandler implements a method for returning, updating and deleting users
func (s Server) UserHandler(c echo.Context) error {
	repo := logics.InitUserRepo(s.Store)
	var req logics.UserReq
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, logics.UserResp{Status: "error", Error: err.Error()})
	}

	//now logic
	switch c.Request().Method {
	case http.MethodGet:
		users, err := logics.GetUsers(repo, req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, logics.UserResp{Status: "error", Error: err.Error()})
		}
		return c.JSON(http.StatusOK, logics.UserResp{Status: "success", Users: users})
	case http.MethodPost:
		err := logics.InsertUsers(repo, req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, logics.UserResp{Status: "error", Error: err.Error()})
		}
		return c.JSON(http.StatusOK, logics.UserResp{Status: "success"})
	case http.MethodPut:
		err := logics.UpdateUsers(repo, req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, logics.UserResp{Status: "error", Error: err.Error()})
		}
	default:
		return c.JSON(http.StatusMethodNotAllowed, logics.UserResp{Status: "error"})
	}
	return nil
}
