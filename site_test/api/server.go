package api

import (
	"context"

	"github.com/labstack/echo/v4"

	"site.test/user"
)

type Users interface {
	Get(ctx context.Context) (loaded []*user.User, err error)
	Create(ctx context.Context, u *user.User) (created *user.User, err error)
}

type Config struct {
	Users Users
}

type Server struct {
	users Users
}

func NewServer(config Config) *Server {
	return &Server{
		users: config.Users,
	}
}

func (server *Server) RegisterTo(e *echo.Echo) {
	e.GET("/users", server.GetUsers)
	e.POST("/users", server.CreateUser)
}
