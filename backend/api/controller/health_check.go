package controller

import (
	"net/http"
	"opeco17/saguru/api/metrics"
	"opeco17/saguru/api/model"
	"time"

	errorsutil "opeco17/saguru/lib/errors"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func HealthCheck(c echo.Context) error {
	// Measure duration
	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	output := new(model.GetHealthCheckOutput)
	output.Status = "healthy"

	if err := c.JSON(http.StatusOK, output); err != nil {
		logrus.Errorf("%#v", errorsutil.Wrap(err, err.Error()))
		return c.String(http.StatusServiceUnavailable, "Something wrong happend")
	}

	return nil
}
