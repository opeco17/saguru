package controller

import (
	"net/http"
	"opeco17/saguru/api/model"
	"opeco17/saguru/api/util"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func GetOrderMetrics(c echo.Context) error {
	logrus.Info("Get order metrics")

	metrics := []string{}
	for _, metric := range util.OrderMetrics() {
		metrics = append(metrics, metric+"_DESC", metric+"_ASC")
	}
	getOrderMetricsOutput := model.GetOrderMetricsOutput{Items: metrics}
	return c.JSON(http.StatusOK, getOrderMetricsOutput)
}
