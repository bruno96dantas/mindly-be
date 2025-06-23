package api

import (
	"database/sql"

	"github.com/bruno96dantas/mindly-be/internal/apps"
	"github.com/bruno96dantas/mindly-be/internal/apps/api/repositories"
	"github.com/bruno96dantas/mindly-be/internal/apps/api/routers"
	"github.com/labstack/echo/v4"
)

type (
	RegisterRouters func(httpServer *echo.Echo, repositories *repositories.Repository)

	API struct {
		Address    string
		HttpServer *echo.Echo
		Repository *repositories.Repository
	}
)

func NewAPI(address, port string, httpServer *echo.Echo, db *sql.DB) apps.Application {
	api := &API{
		Address:    address + ":" + port,
		HttpServer: httpServer,
		Repository: repositories.NewRepository(db),
	}

	api.RegisterRouters(
		routers.HealthcheckGroupRouter,
	)

	return api
}

func (api *API) RegisterRouters(callbacks ...RegisterRouters) {
	for _, callback := range callbacks {
		callback(api.HttpServer, api.Repository)
	}
}

func (api *API) Run() error {
	return api.HttpServer.Start(api.Address)
}
