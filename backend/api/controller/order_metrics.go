package controller

import (
	"net/http"
	"opeco17/saguru/api/metrics"
	"opeco17/saguru/api/model"
	"opeco17/saguru/api/util"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func GetOrderMetrics(c echo.Context) error {
	logrus.Info("Get order metrics")

	since := time.Now()
	defer metrics.M.ObservefunctionCallDuration(since)

	metrics := []string{}
	for _, metric := range util.OrderMetrics() {
		metrics = append(metrics, metric+"_DESC", metric+"_ASC")
	}
	getOrderMetricsOutput := model.GetOrderMetricsOutput{Items: metrics}
	return c.JSON(http.StatusOK, getOrderMetricsOutput)
}
