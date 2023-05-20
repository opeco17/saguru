package main

import (
	"opeco17/saguru/api/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", controller.HealthCheck)
	e.GET("/health", controller.HealthCheck)
	e.GET("/repositories", controller.GetRepositories)
	e.GET("/languages", controller.GetLanguages)
	e.GET("/licenses", controller.GetLicenses)
	e.GET("/labels", controller.GetLabels)
	e.GET("/ordermetrics", controller.GetOrderMetrics)

	e.Logger.Fatal(e.Start(":8000"))
}
