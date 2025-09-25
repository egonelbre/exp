package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"site.test/user"
)

func (server *Server) GetUsers(c echo.Context) error {
	ctx := c.Request().Context()

	users, err := server.users.Get(ctx)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, users)
	return nil
}

func (server *Server) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()

	var u user.User
	if err := c.Bind(&u); err != nil {
		return err
	}

	created, err := server.users.Create(ctx, &u)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, created)
	return nil
}
