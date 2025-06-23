package handlers

import (
	"net/http"

	"github.com/bruno96dantas/mindly-be/internal/apps/api/repositories"
	"github.com/labstack/echo/v4"
)

type (
	HealthcheckHandlerManager struct {
		repository *repositories.Repository
	}
)

func NewHealthcheckHandlerManager(repository *repositories.Repository) *HealthcheckHandlerManager {
	return &HealthcheckHandlerManager{
		repository: repository,
	}
}

func (h *HealthcheckHandlerManager) HealthcheckApp(c echo.Context) error {
	return c.String(http.StatusOK, "WORKING")
}
