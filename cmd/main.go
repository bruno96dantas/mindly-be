package main

import (
	"log/slog"
	"os"

	"github.com/bruno96dantas/mindly-be/internal/apps"
	"github.com/bruno96dantas/mindly-be/internal/apps/api"
	"github.com/bruno96dantas/mindly-be/internal/apps/databases/connection"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := connection.MySQL()
	if err != nil {
		slog.Error("Error connecting to database", "Error", err.Error())
	}

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	api := api.NewAPI(
		"0.0.0.0",
		os.Getenv("PORT"),
		e,
		db,
	)

	if err := apps.Run(api); err != nil {
		slog.Error("main", "error", err.Error())
	}
}
