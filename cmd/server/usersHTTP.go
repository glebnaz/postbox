package server

import (
	"github.com/brianvoe/sjwt"
	"github.com/labstack/echo"
	"net/http"
)

const errUnauthorized = "Bat token, or token is empty"

//middleWare check token
func (s Server) middleWare(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.QueryParam("token")
		verified := sjwt.Verify(token, []byte(s.pass))
		if verified {
			return handlerFunc(c)
		}

		return c.JSON(http.StatusUnauthorized, response{Error: errBadConnectionToStore})
	}
}
