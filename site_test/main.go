package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/sync/errgroup"

	"site.test/api"
	"site.test/formal"
	"site.test/gormdb"
)

func main() {
	err := run(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// start database

	db, err := gormdb.OpenSqlite(ctx, "test.db")
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Migrate(ctx); err != nil {
		return err
	}

	// create services

	formalService := formal.NewService(db.Users())

	// start server

	server := api.NewServer(api.Config{
		Users: formalService,
	})

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	server.RegisterTo(e)

	var group errgroup.Group
	group.Go(func() error {
		<-ctx.Done()
		shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelShutdown()
		return e.Shutdown(shutdownCtx)
	})
	group.Go(func() error {
		defer cancel()
		err := e.Start("127.0.0.1:8000")
		if errors.Is(err, http.ErrServerClosed) {
			err = nil
		}
		return err
	})

	return group.Wait()
}
