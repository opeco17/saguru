package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", Index)
	e.GET("/repositories", GetRepositories)
	e.GET("/languages", GetLanguages)
	e.GET("/licenses", GetLicenses)
	e.GET("/labels", GetLabels)
	e.GET("/ordermetrics", GetOrderMetrics)

	e.Logger.Fatal(e.Start(":8000"))
}
