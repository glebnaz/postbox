package server

import (
	"github.com/brianvoe/sjwt"
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	errBadCred      = "Bad Credentials"
	errUnauthorized = "Bad token, or token is empty"
)

type TokenResponse struct {
	Token string `json:"token,omitempty"`
	Error string `json:"error,omitempty"`
}

//getToken
func (s Server) getToken(c echo.Context) error {
	userQuery := c.Param("user")
	passQuery := c.Param("pass")
	user, pass := s.GetCred()
	if user != userQuery || pass != passQuery {
		return c.JSON(http.StatusUnauthorized, TokenResponse{Error: errBadCred})
	}
	cl := sjwt.New()
	cl.Set("user", user)
	jwt := cl.Generate([]byte(pass))
	return c.JSON(http.StatusOK, TokenResponse{Token: jwt})
}

//middleWare check token
func (s Server) middleWare(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, pass := s.GetCred()
		token := c.QueryParam("token")
		verified := sjwt.Verify(token, []byte(pass))
		if !verified {
			return c.JSON(http.StatusUnauthorized, response{Error: errUnauthorized})
		}
		cl, err := sjwt.Parse(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, response{Error: errUnauthorized})
		}
		userFromToken, err := cl.Get("user")
		if userFromToken != user || err != nil {
			return c.JSON(http.StatusUnauthorized, response{Error: errUnauthorized})
		}
		return handlerFunc(c)
	}
}
