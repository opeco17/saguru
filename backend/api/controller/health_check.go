package controller

import (
	"net/http"
	"opeco17/saguru/api/metrics"
	"opeco17/saguru/api/model"
	"time"

	"github.com/labstack/echo/v4"
)

func HealthCheck(c echo.Context) error {
	// Measure duration
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	output := new(model.GetHealthCheckOutput)
	output.Status = "healthy"

	return c.JSON(http.StatusOK, output)
}
