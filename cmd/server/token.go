package server

import (
	"net/http"

	"github.com/glebnaz/postbox/internal/errors"

	"github.com/glebnaz/postbox/internal/secure"
	"github.com/labstack/echo/v4"
)

type tokenRequest struct {
	User string `query:"user"`
	Pass string `query:"pass"`
}

type tokenResponse struct {
	Token string `json:"token,omitempty"`
	Error string `json:"error,omitempty"`
}

//getToken
func (s Server) getToken(c echo.Context) error {
	var req tokenRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, tokenResponse{Error: errors.Unauthorized.New(err).Error()})
	}
	user, pass := s.GetCred()
	if user != req.User || pass != req.Pass {
		return c.JSON(http.StatusUnauthorized, tokenResponse{Error: errors.BadCred.Error()})
	}
	jwt := secure.GenerateJWT(user, pass)
	return c.JSON(http.StatusOK, tokenResponse{Token: jwt})
}

//middleWare check token
func (s Server) middleWare(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, pass := s.GetCred()
		token := c.QueryParam("token")
		if !secure.ValidateJWT(token, user, pass) {
			return c.JSON(http.StatusUnauthorized, tokenResponse{Error: errors.Unauthorized.Error()})
		}
		return handlerFunc(c)
	}
}
