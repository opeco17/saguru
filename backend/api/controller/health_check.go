package controller

import (
	"net/http"
	"opeco17/saguru/api/model"

	"github.com/labstack/echo/v4"
)

func HealthCheck(c echo.Context) error {
	output := new(model.GetHealthCheckOutput)
	output.Status = "healthy"
	return c.JSON(http.StatusOK, output)
}
