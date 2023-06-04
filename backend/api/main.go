package main

import (
	"fmt"
	"net/http"
	"opeco17/saguru/api/controller"
	"opeco17/saguru/api/metrics"
	"opeco17/saguru/api/util"
	"opeco17/saguru/lib/memcached"

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
	e.GET("/test", func(c echo.Context) error {
		client, _ := util.GetMemcachedClient()
		licenses, _ := memcached.GetLicenses(client)
		fmt.Println("test")
		return c.JSON(http.StatusOK, licenses)
	})

	e.Logger.Fatal(e.Start(":8000"))
}
