package main

import (
	"opeco17/saguru/api/controller"
	"opeco17/saguru/api/metrics"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	p := prometheus.NewPrometheus("echo", nil, metrics.M.MetricList())
	p.Use(e)

	e.GET("/", controller.HealthCheck)
	e.GET("/health", controller.HealthCheck)
	e.GET("/repositories", controller.GetRepositories)
	e.GET("/languages", controller.GetLanguages)
	e.GET("/licenses", controller.GetLicenses)
	e.GET("/labels", controller.GetLabels)
	e.GET("/ordermetrics", controller.GetOrderMetrics)

	e.Logger.Fatal(e.Start(":8000"))
}
