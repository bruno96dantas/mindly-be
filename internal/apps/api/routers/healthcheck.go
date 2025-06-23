package routers

import (
	"github.com/bruno96dantas/mindly-be/internal/apps/api/handlers"
	"github.com/bruno96dantas/mindly-be/internal/apps/api/repositories"
	"github.com/labstack/echo/v4"
)

func HealthcheckGroupRouter(mux *echo.Echo, repositories *repositories.Repository) {
	group := mux.Group("/healthcheck")

	healthcheckHandlers := handlers.NewHealthcheckHandlerManager(repositories)

	group.GET("", healthcheckHandlers.HealthcheckApp)
}
